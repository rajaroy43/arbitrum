/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <avm/machine.hpp>

#include <avm/opcodes.hpp>
#include <avm/util.hpp>

#include <iostream>

namespace {
std::vector<CodePoint> opsToCodePoints(const std::vector<Operation>& ops) {
    std::vector<CodePoint> cps;
    cps.reserve(ops.size());
    uint64_t pc = 0;
    for (auto& op : ops) {
        cps.emplace_back(pc, std::move(op), 0);
        pc++;
    }
    for (uint64_t i = 0; i < cps.size() - 1; i++) {
        cps[cps.size() - 2 - i].nextHash = hash(cps[cps.size() - 1 - i]);
    }
    return cps;
}
}  // namespace

class bad_pop_type : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "bad_variant_access";
    }
};

class int_out_of_bounds : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "int_out_of_bounds";
    }
};

MachineState::MachineState()
    : pool(std::make_unique<TuplePool>()), context({0, 0}) {}

uint256_t MachineState::hash() const {
    if (state == Status::Halted)
        return 0;
    if (state == Status::Error)
        return 1;

    std::array<unsigned char, 32 * 6> data;
    auto oit = data.begin();
    {
        auto val = ::hash(code[pc]);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = stack.hash();
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = auxstack.hash();
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(registerVal);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(staticVal);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(errpc);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(data.data(), 32 * 6, hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

std::ostream& operator<<(std::ostream& os, const MachineState& val) {
    os << "codePointHash " << to_hex_str(hash(val.code[val.pc])) << "\n";
    os << "stackHash " << to_hex_str(val.stack.hash()) << "\n";
    os << "auxStackHash " << to_hex_str(val.auxstack.hash()) << "\n";
    os << "registerHash " << to_hex_str(hash(val.registerVal)) << "\n";
    os << "staticHash " << to_hex_str(hash(val.staticVal)) << "\n";
    os << "errHandlerHash " << to_hex_str(hash(val.errpc)) << "\n";
    return os;
}

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.m;
    return os;
}

uint256_t& assumeInt(value& val) {
    auto aNum = nonstd::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

const uint256_t& assumeInt(const value& val) {
    auto aNum = nonstd::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

uint64_t assumeInt64(uint256_t& val) {
    if (val > std::numeric_limits<uint64_t>::max())
        throw int_out_of_bounds{};

    return static_cast<uint64_t>(val);
}

Tuple& assumeTuple(value& val) {
    auto tup = nonstd::get_if<Tuple>(&val);
    if (!tup) {
        throw bad_pop_type{};
    }
    return *tup;
}

void MachineState::deserialize(char* bufptr) {
    uint32_t version;
    memcpy(&version, bufptr, sizeof(version));
    version = __builtin_bswap32(version);
    bufptr += sizeof(version);

    if (version != CURRENT_AO_VERSION) {
        std::cout << "incorrect version of .ao file" << std::endl;
        std::cout << "expected version " << CURRENT_AO_VERSION
                  << " found version " << version << std::endl;
        return;
    }

    uint32_t extentionId = 1;
    while (extentionId != 0) {
        memcpy(&extentionId, bufptr, sizeof(extentionId));
        extentionId = __builtin_bswap32(extentionId);
        bufptr += sizeof(extentionId);
        if (extentionId > 0) {
            std::cout << "found extention" << std::endl;
        }
    }
    uint64_t codeCount;
    memcpy(&codeCount, bufptr, sizeof(codeCount));
    bufptr += sizeof(codeCount);
    codeCount = boost::endian::big_to_native(codeCount);
    code.reserve(codeCount);

    std::vector<Operation> ops;
    for (uint64_t i = 0; i < codeCount; i++) {
        ops.emplace_back(deserializeOperation(bufptr, *pool));
    }
    code = opsToCodePoints(ops);

    staticVal = deserialize_value(bufptr, *pool);
    pc = 0;
}

bool MachineState::hasPendingMessages() const {
    return !(pendingInbox == Tuple());
}

void MachineState::sendOnchainMessage(const Message& msg) {
    pendingInbox = Tuple{uint256_t{0}, std::move(pendingInbox),
                         msg.toValue(*pool), pool.get()};
    balance.add(msg.token, msg.currency);
}

void MachineState::deliverMessageStack(Tuple&& messages) {
    if (!(messages == Tuple())) {
        inbox = Tuple(uint256_t(1), std::move(inbox), std::move(messages),
                      pool.get());
    }
}

void MachineState::sendOffchainMessages(const std::vector<Message>& messages) {
    Tuple messageStack;
    for (const auto& message : messages) {
        auto messageVal = message.toValue(*pool);
        messageStack = Tuple{uint256_t{0}, std::move(messageStack),
                             std::move(messageVal), pool.get()};
    }
    deliverMessageStack(std::move(messageStack));
}

void MachineState::deliverOnchainMessages() {
    deliverMessageStack(std::move(pendingInbox));
    pendingInbox = Tuple();
}

void uint256_t_to_buf(uint256_t val, std::vector<unsigned char>& buf) {
    std::vector<unsigned char> tmpbuf;
    tmpbuf.resize(32);
    to_big_endian(val, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

std::vector<unsigned char> MachineState::marshalForProof() {
    std::vector<unsigned char> buf;
    std::vector<bool> stackPops = InstructionStackPops.at(code[pc].op.opcode);
    if (code[pc].op.immediate) {
        stackPops.erase(stackPops.begin());
    }
    std::vector<bool> auxStackPops =
        InstructionAuxStackPops.at(code[pc].op.opcode);
    std::vector<value> stackVals;
    uint256_t baseStackHash = stack.SolidityProofValue(stackPops, stackVals);
    std::vector<value> auxStackVals;
    uint256_t baseAuxStackHash =
        auxstack.SolidityProofValue(auxStackPops, auxStackVals);
    uint256_t registerHash = ::hash(registerVal);
    uint256_t staticHash = ::hash(staticVal);
    uint256_t errHandlerHash = ::hash(errpc);
    std::cout << "Proof of " << code[pc] << " has " << stackVals.size()
              << " stack vals and " << auxStackVals.size() << " aux stack vals"
              << std::endl;
    std::cout << "pc next hash " << to_hex_str(code[pc].nextHash) << std::endl;
    std::cout << "baseStackHash " << to_hex_str(baseStackHash) << std::endl;
    std::cout << "baseAuxStackHash " << to_hex_str(baseAuxStackHash)
              << std::endl;
    std::cout << "registerHash " << to_hex_str(registerHash) << std::endl;
    std::cout << "staticHash " << to_hex_str(staticHash) << std::endl;
    std::cout << "errHandlerHash " << to_hex_str(errHandlerHash) << std::endl;
    uint256_t_to_buf(code[pc].nextHash, buf);
    uint256_t_to_buf(baseStackHash, buf);
    uint256_t_to_buf(baseAuxStackHash, buf);
    uint256_t_to_buf(registerHash, buf);
    uint256_t_to_buf(staticHash, buf);
    uint256_t_to_buf(errHandlerHash, buf);
    code[pc].op.marshal(buf);
    for (auto const& stackval : stackVals) {
        marshal_value(stackval, buf);
    }
    for (auto const& auxstackval : auxStackVals) {
        marshal_value(auxstackval, buf);
    }
    std::cout << "marshal size " << buf.size() << std::endl;
    return buf;
}

void Machine::sendOnchainMessage(const Message& msg) {
    m.sendOnchainMessage(msg);
}

void Machine::deliverOnchainMessages() {
    m.deliverOnchainMessages();
}

void Machine::sendOffchainMessages(const std::vector<Message>& messages) {
    m.sendOffchainMessages(messages);
}

Assertion Machine::run(uint64_t stepCount,
                       uint64_t timeBoundStart,
                       uint64_t timeBoundEnd) {
    m.context = AssertionContext{TimeBounds{{timeBoundStart, timeBoundEnd}}};
    uint64_t i;
    if (m.state == Status::Blocked) {
        m.state = Status::Extensive;
    }
    for (i = 0; i < stepCount; i++) {
        if (m.state == Status::Error || m.state == Status::Halted ||
            m.state == Status::Blocked) {
            break;
        }
        auto ret = runOne();
        if (ret < 0) {
            break;
        }
    }

    if (m.state == Status::Blocked) {
        i--;
    }

    if (m.state == Status::Error) {
        std::cout << "error state" << std::endl;
        if (m.errpc.isSet()) {
            m.pc = m.errpc.pc;
        }
    }
    if (m.state == Status::Halted) {
        // set error return
        //        std::cout << "halted state" << std::endl;
    }
    return {i, std::move(m.context.outMessage), std::move(m.context.logs)};
}

int Machine::runOne() {
    //    std::cout << to_hex_str(hash()) << " " << m.code[m.pc].op <<
    //    std::endl; std::cout << *this << std::endl; std::cout<<"in
    //    runOne"<<std::endl;
    if (m.state == Status::Error) {
        // set error return
        std::cout << "error state" << std::endl;
        return -1;
    }

    if (m.state == Status::Halted) {
        // set error return
        std::cout << "halted state" << std::endl;
        std::cout << "full stack - size=" << m.stack.stacksize() << std::endl;
        while (m.stack.stacksize() > 0) {
            std::cout << m.stack[0] << std::endl;
            m.stack.popClear();
        }
        return -2;
    }

    auto& instruction = m.code[m.pc];
    //    std::cout<<m.pc<<"
    //    "<<InstructionNames.at(instruction.op.opcode)<<std::endl; std::cout <<
    //    to_hex_str(m.hash()) << "\n" << m << std::endl; std::cout << m <<
    //    std::endl;
    if (instruction.op.immediate) {
        //        std::cout<<"immediateVal = "<<*immediateVal<<std::endl;
        auto imm = *instruction.op.immediate;
        m.stack.push(std::move(imm));
    }

    try {
        //        std::cout<<"calling runInstruction"<<std::endl;
        m.runOp(instruction.op.opcode);
        //        std::cout<<"after runInstruction stack size=
        //        "<<stack.stacksize()<< std::endl; if (stack.stacksize()>0){
        //            std::cout<<"top="<<stack.peek()<< std::endl;
        //        }
    } catch (const bad_pop_type& e) {
        m.state = Status::Error;
    } catch (const bad_tuple_index& e) {
        m.state = Status::Error;
    }

    return 0;
}

template <typename T>
static T shrink(uint256_t i) {
    return static_cast<T>(i & std::numeric_limits<T>::max());
}

namespace {
static void add(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum + bNum;
    m.stack.popClear();
    ++m.pc;
}

static void mul(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum * bNum;
    m.stack.popClear();
    ++m.pc;
}

static void sub(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum - bNum;
    m.stack.popClear();
    ++m.pc;
}

static void div(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = aNum / bNum;
    }
    m.stack.popClear();
    ++m.pc;
}

static void sdiv(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    const auto min = (std::numeric_limits<uint256_t>::max() / 2) + 1;

    if (bNum == 0) {
        m.state = Status::Error;
    } else if (aNum == min && bNum == -1) {
        m.stack[1] = aNum;
    } else {
        const auto signA = get_sign(aNum);
        const auto signB = get_sign(bNum);
        if (signA == -1)
            aNum = 0 - aNum;
        if (signB == -1)
            bNum = 0 - bNum;
        m.stack[1] = (aNum / bNum) * signA * signB;
    }
    m.stack.popClear();
    ++m.pc;
}

static void mod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum != 0) {
        m.stack[1] = aNum % bNum;
    } else {
        m.state = Status::Error;
    }
    m.stack.popClear();
    ++m.pc;
}

static void smod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        const auto signA = get_sign(aNum);
        const auto signB = get_sign(bNum);
        if (signA == -1)
            aNum = 0 - aNum;
        if (signB == -1)
            bNum = 0 - bNum;
        m.stack[1] = (aNum % bNum) * signA;
    }
    m.stack.popClear();
    ++m.pc;
}

static void addmod(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    auto& cNum = assumeInt(m.stack[2]);

    if (cNum == 0) {
        m.state = Status::Error;
    } else {
        uint512_t aBig = aNum;
        uint512_t bBig = bNum;
        m.stack[2] = static_cast<uint256_t>((aBig + bBig) % cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

static void mulmod(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    auto& cNum = assumeInt(m.stack[2]);

    if (cNum == 0) {
        m.state = Status::Error;
    } else {
        uint512_t aBig = aNum;
        uint512_t bBig = bNum;
        m.stack[2] = static_cast<uint256_t>((aBig * bBig) % cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

static void exp(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    uint64_t bSmall = assumeInt64(bNum);
    m.stack[1] = power(aNum, bSmall);
    m.stack.popClear();
    ++m.pc;
}

static void lt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum < bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

static void gt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum > bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

static void slt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (aNum == bNum) {
        m.stack[1] = 0;
    } else {
        uint8_t signA = get_sign(aNum);
        uint8_t signB = get_sign(bNum);

        if (signA != signB) {
            m.stack[1] = signA == 1 ? 0 : 1;
        } else {
            m.stack[1] = aNum < bNum ? 1 : 0;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

static void sgt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (aNum == bNum) {
        m.stack[1] = 0;
    } else {
        uint8_t signA = get_sign(aNum);
        uint8_t signB = get_sign(bNum);

        if (signA != signB) {
            m.stack[1] = signA == 1 ? 1 : 0;
        } else {
            m.stack[1] = aNum > bNum ? 1 : 0;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

static void eq(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aVal = m.stack[0];
    auto& bVal = m.stack[1];
    m.stack[1] = aVal == bVal ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

static void iszero(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = aNum.is_zero() ? 1 : 0;
    ++m.pc;
}

static void bitwiseAnd(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum & bNum;
    m.stack.popClear();
    ++m.pc;
}

static void bitwiseOr(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum | bNum;
    m.stack.popClear();
    ++m.pc;
}

static void bitwiseXor(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum ^ bNum;
    m.stack.popClear();
    ++m.pc;
}

static void bitwiseNot(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = ~aNum;
    ++m.pc;
}

static void byte(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum >= 32) {
        m.stack[1] = 0;
    } else {
        const auto shift = 256 - 8 - 8 * shrink<uint8_t>(bNum);
        const auto mask = uint256_t(255) << shift;
        m.stack[1] = (aNum & mask) >> shift;
    }
    m.stack.popClear();
    ++m.pc;
}

static void signExtend(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum >= 32) {
        m.stack[1] = m.stack[0];
    } else {
        uint256_t t = 248 - 8 * bNum;
        int signBit = bit(aNum, (int)(255 - t));
        uint256_t mask = power(uint256_t(2), uint64_t(255 - t)) - 1;
        if (signBit == 0) {
            m.stack[1] = aNum & mask;
        } else {
            mask ^= -1;
            m.stack[1] = aNum | mask;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

static void hashOp(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = ::hash(m.stack[0]);
    ++m.pc;
}

static void typeOp(MachineState& m) {
    m.stack.prepForMod(1);
    if (nonstd::holds_alternative<uint256_t>(m.stack[0]))
        m.stack[0] = NUM;
    else if (nonstd::holds_alternative<CodePoint>(m.stack[0]))
        m.stack[0] = CODEPT;
    else if (nonstd::holds_alternative<Tuple>(m.stack[0]))
        m.stack[0] = TUPLE;
    ++m.pc;
}

static void pop(MachineState& m) {
    m.stack.popClear();
    ++m.pc;
}

static void spush(MachineState& m) {
    value copiedStatic = m.staticVal;
    m.stack.push(std::move(copiedStatic));
    ++m.pc;
}

static void rpush(MachineState& m) {
    value copiedRegister = m.registerVal;
    m.stack.push(std::move(copiedRegister));
    ++m.pc;
}

static void rset(MachineState& m) {
    m.stack.prepForMod(1);
    m.registerVal = m.stack[0];
    m.stack.popClear();
    ++m.pc;
}

static void jump(MachineState& m) {
    m.stack.prepForMod(1);
    auto target = nonstd::get_if<CodePoint>(&m.stack[0]);
    if (target) {
        m.pc = target->pc;
    } else {
        m.state = Status::Error;
    }
    m.stack.popClear();
}

static void cjump(MachineState& m) {
    m.stack.prepForMod(2);
    auto target = nonstd::get_if<CodePoint>(&m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum != 0) {
        if (target) {
            m.pc = target->pc;
        } else {
            m.state = Status::Error;
        }
    } else {
        ++m.pc;
    }
    m.stack.popClear();
    m.stack.popClear();
}

static void stackEmpty(MachineState& m) {
    if (m.stack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

static void pcPush(MachineState& m) {
    m.stack.push(m.code[m.pc]);
    ++m.pc;
}

static void auxPush(MachineState& m) {
    m.stack.prepForMod(1);
    m.auxstack.push(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

static void auxPop(MachineState& m) {
    m.auxstack.prepForMod(1);
    m.stack.push(std::move(m.auxstack[0]));
    m.auxstack.popClear();
    ++m.pc;
}

static void auxStackEmpty(MachineState& m) {
    if (m.auxstack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

static void errPush(MachineState& m) {
    m.stack.push(m.errpc);
    ++m.pc;
}

static void errSet(MachineState& m) {
    m.stack.prepForMod(1);
    auto codePointVal = nonstd::get_if<CodePoint>(&m.stack[0]);
    if (!codePointVal) {
        m.state = Status::Error;
    } else {
        m.errpc = *codePointVal;
    }
    m.stack.popClear();
    ++m.pc;
}

static void dup0(MachineState& m) {
    value valACopy = m.stack[0];
    m.stack.push(std::move(valACopy));
    ++m.pc;
}

static void dup1(MachineState& m) {
    value valBCopy = m.stack[1];
    m.stack.push(std::move(valBCopy));
    ++m.pc;
}

static void dup2(MachineState& m) {
    value valCCopy = m.stack[2];
    m.stack.push(std::move(valCCopy));
    ++m.pc;
}

static void swap1(MachineState& m) {
    m.stack.prepForMod(2);
    value temp = m.stack[0];
    m.stack[0] = m.stack[1];
    m.stack[1] = temp;
    ++m.pc;
}

static void swap2(MachineState& m) {
    m.stack.prepForMod(3);
    value temp = m.stack[0];
    m.stack[0] = m.stack[2];
    m.stack[2] = temp;
    ++m.pc;
}

static void tget(MachineState& m) {
    m.stack.prepForMod(2);
    auto& index = assumeInt(m.stack[0]);
    auto& tup = assumeTuple(m.stack[1]);
    m.stack[1] = tup.get_element(static_cast<uint32_t>(index));
    m.stack.popClear();
    ++m.pc;
}

static void tset(MachineState& m) {
    m.stack.prepForMod(3);
    auto& index = assumeInt(m.stack[0]);
    auto& tup = assumeTuple(m.stack[1]);
    tup.set_element(static_cast<uint32_t>(index), std::move(m.stack[2]));
    m.stack[2] = std::move(tup);
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

static void tlen(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = assumeTuple(m.stack[0]).tuple_size();
    ++m.pc;
}

static void breakpoint(MachineState& m) {
    m.state = Status::Blocked;
}

static void log(MachineState& m) {
    m.stack.prepForMod(1);
    m.context.logs.push_back(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

static void debug(MachineState& m) {
    datastack tmpstk;
    std::cout << std::endl;
    std::cout << "full stack - size=" << m.stack.stacksize() << std::endl;
    while (m.stack.stacksize() > 0) {
        std::cout << m.stack[0] << std::endl;
        tmpstk.push(std::move(m.stack[0]));
        m.stack.popClear();
    }
    while (tmpstk.stacksize() > 0) {
        m.stack.push(std::move(tmpstk[0]));
        tmpstk.popClear();
    }
    std::cout << "register val=" << m.registerVal << std::endl << std::endl;
    ++m.pc;
}

static void send(MachineState& m) {
    m.stack.prepForMod(1);
    Message outMsg;
    auto success = outMsg.deserialize(m.stack[0]);
    if (!success) {
        m.state = Status::Error;
        return;
    }
    if (!m.balance.Spend(outMsg.token, outMsg.currency)) {
        m.state = Status::Blocked;
    } else {
        m.stack.popClear();
        m.context.outMessage.push_back(outMsg);
        ++m.pc;
    }
}

static void nbsend(MachineState& m) {
    m.stack.prepForMod(1);

    Message outMsg;
    auto success = outMsg.deserialize(m.stack[0]);
    if (!success) {
        m.state = Status::Error;
        return;
    }

    bool spent = m.balance.Spend(outMsg.token, outMsg.currency);
    if (!spent) {
        m.stack[0] = 0;
    } else {
        m.context.outMessage.push_back(outMsg);
        m.stack[0] = 1;
    }
    ++m.pc;
}

static void getTime(MachineState& m) {
    Tuple tup(m.pool.get(), 2);
    tup.set_element(0, m.context.timeBounds[0]);
    tup.set_element(1, m.context.timeBounds[1]);
    m.stack.push(std::move(tup));
    ++m.pc;
}

static void inboxOp(MachineState& m) {
    m.stack.prepForMod(1);
    auto stackTop = nonstd::get_if<Tuple>(&m.stack[0]);
    if (stackTop && m.inbox == *stackTop) {
        m.state = Status::Blocked;
    } else {
        value inboxCopy = m.inbox;
        m.stack[0] = std::move(inboxCopy);
        ++m.pc;
    }
}
}  // namespace

void MachineState::runOp(OpCode opcode) {
    switch (opcode) {
        /**************************/
        /*  Arithmetic Operations */
        /**************************/
        case OpCode::ADD:
            add(*this);
            break;
        case OpCode::MUL:
            mul(*this);
            break;
        case OpCode::SUB:
            sub(*this);
            break;
        case OpCode::DIV:
            div(*this);
            break;
        case OpCode::SDIV:
            sdiv(*this);
            break;
        case OpCode::MOD:
            mod(*this);
            break;
        case OpCode::SMOD:
            smod(*this);
            break;
        case OpCode::ADDMOD:
            addmod(*this);
            break;
        case OpCode::MULMOD:
            mulmod(*this);
            break;
        case OpCode::EXP:
            exp(*this);
            break;
        /******************************************/
        /*  Comparison & Bitwise Logic Operations */
        /******************************************/
        case OpCode::LT:
            lt(*this);
            break;
        case OpCode::GT:
            gt(*this);
            break;
        case OpCode::SLT:
            slt(*this);
            break;
        case OpCode::SGT:
            sgt(*this);
            break;
        case OpCode::EQ:
            eq(*this);
            break;
        case OpCode::ISZERO:
            iszero(*this);
            break;
        case OpCode::BITWISE_AND:
            bitwiseAnd(*this);
            break;
        case OpCode::BITWISE_OR:
            bitwiseOr(*this);
            break;
        case OpCode::BITWISE_XOR:
            bitwiseXor(*this);
            break;
        case OpCode::BITWISE_NOT:
            bitwiseNot(*this);
            break;
        case OpCode::BYTE:
            byte(*this);
            break;
        case OpCode::SIGNEXTEND:
            signExtend(*this);
            break;

        /***********************/
        /*  Hashing Operations */
        /***********************/
        case OpCode::HASH:
            hashOp(*this);
            break;

        case OpCode::TYPE:
            typeOp(*this);
            break;

        /***********************************************/
        /*  Stack, Memory, Storage and Flow Operations */
        /***********************************************/
        case OpCode::POP:
            pop(*this);
            break;
        case OpCode::SPUSH:
            spush(*this);
            break;
        case OpCode::RPUSH:
            rpush(*this);
            break;
        case OpCode::RSET:
            rset(*this);
            break;
        case OpCode::JUMP:
            jump(*this);
            break;
        case OpCode::CJUMP:
            cjump(*this);
            break;
        case OpCode::STACKEMPTY:
            stackEmpty(*this);
            break;
        case OpCode::PCPUSH:
            pcPush(*this);
            break;
        case OpCode::AUXPUSH:
            auxPush(*this);
            break;
        case OpCode::AUXPOP:
            auxPop(*this);
            break;
        case OpCode::AUXSTACKEMPTY:
            auxStackEmpty(*this);
            break;
        case OpCode::NOP:
            ++pc;
            break;
        case OpCode::ERRPUSH:
            errPush(*this);
            break;
        case OpCode::ERRSET:
            errSet(*this);
            break;
            /****************************************/
            /*  Duplication and Exchange Operations */
            /****************************************/
        case OpCode::DUP0:
            dup0(*this);
            break;
        case OpCode::DUP1:
            dup1(*this);
            break;
        case OpCode::DUP2:
            dup2(*this);
            break;
        case OpCode::SWAP1:
            swap1(*this);
            break;
        case OpCode::SWAP2:
            swap2(*this);
            break;
            /*********************/
            /*  Tuple Operations */
            /*********************/
        case OpCode::TGET:
            tget(*this);
            break;
        case OpCode::TSET:
            tset(*this);
            break;
        case OpCode::TLEN:
            tlen(*this);
            break;
            /***********************/
            /*  Logging Operations */
            /***********************/
        case OpCode::BREAKPOINT:
            breakpoint(*this);
            break;
        case OpCode::LOG:
            log(*this);
            break;
        case OpCode::DEBUG:
            debug(*this);
            break;
            /**********************/
            /*  System Operations */
            /**********************/
        case OpCode::SEND:
            send(*this);
            break;
        case OpCode::NBSEND:
            nbsend(*this);
            break;
        case OpCode::GETTIME:
            getTime(*this);
            break;
        case OpCode::INBOX:
            inboxOp(*this);
            break;
        case OpCode::ERROR:
            state = Status::Error;
            if (errpc.isSet()) {
                pc = errpc.pc;
            }
            break;
        case OpCode::HALT:
            std::cout << "Hit halt opcode at instruction " << pc << "\n";
            state = Status::Halted;
            break;
        default:
            std::stringstream ss;
            ss << "Unhandled opcode <" << InstructionNames.at(opcode) << ">"
               << std::hex << static_cast<int>(opcode);
            throw std::runtime_error(ss.str());
    }
}