/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer } from 'ethers'
import { Provider, TransactionRequest } from '@ethersproject/providers'
import { Contract, ContractFactory, Overrides } from '@ethersproject/contracts'

import type { Cloneable } from '../Cloneable'

export class Cloneable__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(overrides?: Overrides): Promise<Cloneable> {
    return super.deploy(overrides || {}) as Promise<Cloneable>
  }
  getDeployTransaction(overrides?: Overrides): TransactionRequest {
    return super.getDeployTransaction(overrides || {})
  }
  attach(address: string): Cloneable {
    return super.attach(address) as Cloneable
  }
  connect(signer: Signer): Cloneable__factory {
    return super.connect(signer) as Cloneable__factory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): Cloneable {
    return new Contract(address, _abi, signerOrProvider) as Cloneable
  }
}

const _abi = [
  {
    inputs: [],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  {
    inputs: [],
    name: 'isMaster',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
]

const _bytecode =
  '0x6080604052348015600f57600080fd5b506000805460ff1916600117905560868061002b6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80636f791d2914602d575b600080fd5b60336047565b604080519115158252519081900360200190f35b60005460ff169056fea264697066735822122058b9ee7254c959b6503b4633dc6b8c4f765d5ee87fdf5bbcc9e718528d29f83a64736f6c634300060b0033'
