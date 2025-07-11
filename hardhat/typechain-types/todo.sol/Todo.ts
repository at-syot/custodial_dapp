/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedListener,
  TypedContractMethod,
} from "../common";

export declare namespace Todo {
  export type TaskStruct = { content: string; status: boolean };

  export type TaskStructOutput = [content: string, status: boolean] & {
    content: string;
    status: boolean;
  };
}

export interface TodoInterface extends Interface {
  getFunction(
    nameOrSignature: "add" | "get" | "list" | "remove" | "update"
  ): FunctionFragment;

  encodeFunctionData(functionFragment: "add", values: [string]): string;
  encodeFunctionData(functionFragment: "get", values: [BigNumberish]): string;
  encodeFunctionData(functionFragment: "list", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "remove",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "update",
    values: [BigNumberish, string, boolean]
  ): string;

  decodeFunctionResult(functionFragment: "add", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "get", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "list", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "remove", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "update", data: BytesLike): Result;
}

export interface Todo extends BaseContract {
  connect(runner?: ContractRunner | null): Todo;
  waitForDeployment(): Promise<this>;

  interface: TodoInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  add: TypedContractMethod<[content: string], [void], "nonpayable">;

  get: TypedContractMethod<[id: BigNumberish], [Todo.TaskStructOutput], "view">;

  list: TypedContractMethod<[], [Todo.TaskStructOutput[]], "view">;

  remove: TypedContractMethod<[id: BigNumberish], [void], "nonpayable">;

  update: TypedContractMethod<
    [id: BigNumberish, content: string, status: boolean],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "add"
  ): TypedContractMethod<[content: string], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "get"
  ): TypedContractMethod<[id: BigNumberish], [Todo.TaskStructOutput], "view">;
  getFunction(
    nameOrSignature: "list"
  ): TypedContractMethod<[], [Todo.TaskStructOutput[]], "view">;
  getFunction(
    nameOrSignature: "remove"
  ): TypedContractMethod<[id: BigNumberish], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "update"
  ): TypedContractMethod<
    [id: BigNumberish, content: string, status: boolean],
    [void],
    "nonpayable"
  >;

  filters: {};
}
