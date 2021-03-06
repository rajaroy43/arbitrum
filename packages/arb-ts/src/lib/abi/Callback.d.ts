/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import {
  ethers,
  EventFilter,
  Signer,
  BigNumber,
  BigNumberish,
  PopulatedTransaction,
} from 'ethers'
import {
  Contract,
  ContractTransaction,
  Overrides,
  CallOverrides,
} from '@ethersproject/contracts'
import { BytesLike } from '@ethersproject/bytes'
import { Listener, Provider } from '@ethersproject/providers'
import { FunctionFragment, EventFragment, Result } from '@ethersproject/abi'

interface CallbackInterface extends ethers.utils.Interface {
  functions: {
    'doCallback()': FunctionFragment
    'sendDummies()': FunctionFragment
  }

  encodeFunctionData(functionFragment: 'doCallback', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'sendDummies',
    values?: undefined
  ): string

  decodeFunctionResult(functionFragment: 'doCallback', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'sendDummies', data: BytesLike): Result

  events: {
    'DummyEvent(uint256,uint256,uint256)': EventFragment
  }

  getEvent(nameOrSignatureOrTopic: 'DummyEvent'): EventFragment
}

export class Callback extends Contract {
  connect(signerOrProvider: Signer | Provider | string): this
  attach(addressOrName: string): this
  deployed(): Promise<this>

  on(event: EventFilter | string, listener: Listener): this
  once(event: EventFilter | string, listener: Listener): this
  addListener(eventName: EventFilter | string, listener: Listener): this
  removeAllListeners(eventName: EventFilter | string): this
  removeListener(eventName: any, listener: Listener): this

  interface: CallbackInterface

  functions: {
    doCallback(overrides?: Overrides): Promise<ContractTransaction>

    'doCallback()'(overrides?: Overrides): Promise<ContractTransaction>

    sendDummies(overrides?: Overrides): Promise<ContractTransaction>

    'sendDummies()'(overrides?: Overrides): Promise<ContractTransaction>
  }

  doCallback(overrides?: Overrides): Promise<ContractTransaction>

  'doCallback()'(overrides?: Overrides): Promise<ContractTransaction>

  sendDummies(overrides?: Overrides): Promise<ContractTransaction>

  'sendDummies()'(overrides?: Overrides): Promise<ContractTransaction>

  callStatic: {
    doCallback(overrides?: CallOverrides): Promise<[BigNumber, BigNumber]>

    'doCallback()'(overrides?: CallOverrides): Promise<[BigNumber, BigNumber]>

    sendDummies(overrides?: CallOverrides): Promise<void>

    'sendDummies()'(overrides?: CallOverrides): Promise<void>
  }

  filters: {
    DummyEvent(a: BigNumberish | null, b: null, c: null): EventFilter
  }

  estimateGas: {
    doCallback(overrides?: Overrides): Promise<BigNumber>

    'doCallback()'(overrides?: Overrides): Promise<BigNumber>

    sendDummies(overrides?: Overrides): Promise<BigNumber>

    'sendDummies()'(overrides?: Overrides): Promise<BigNumber>
  }

  populateTransaction: {
    doCallback(overrides?: Overrides): Promise<PopulatedTransaction>

    'doCallback()'(overrides?: Overrides): Promise<PopulatedTransaction>

    sendDummies(overrides?: Overrides): Promise<PopulatedTransaction>

    'sendDummies()'(overrides?: Overrides): Promise<PopulatedTransaction>
  }
}
