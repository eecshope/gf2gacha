// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';

export function GetGameInfo():Promise<model.Info>;

export function GetPoolInfo(arg1:string,arg2:number):Promise<model.Pool>;

export function GetUserList():Promise<Array<string>>;

export function IncrementalUpdatePoolInfo():Promise<string>;
