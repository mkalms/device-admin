import {
  GetDeviceConfigResponseDuplexEnum,
  SetDeviceConfigRequest,
} from "../generated/api";
import { api } from "../backendApi";

export interface DeviceConfigState {
  ipAddress: string;
  netMask: string;
  bitRate: number;
  duplex: GetDeviceConfigResponseDuplexEnum;
}

export async function getDeviceConfig(): Promise<DeviceConfigState> {
  const getDeviceConfigResponse = await api.getDeviceConfig();
  const deviceConfigState = {
    ipAddress: getDeviceConfigResponse.data.ipAddress,
    netMask: getDeviceConfigResponse.data.netMask,
    bitRate: getDeviceConfigResponse.data.bitRate,
    duplex: getDeviceConfigResponse.data.duplex,
  } as DeviceConfigState;

  return deviceConfigState;
}

export async function setDeviceConfig(newConfig: DeviceConfigState) {
  const setDeviceConfigRequest = {
    ipAddress: newConfig.ipAddress,
    netMask: newConfig.netMask,
    bitRate: newConfig.bitRate,
    duplex: newConfig.duplex,
  } as SetDeviceConfigRequest;
  await api.setDeviceConfig(setDeviceConfigRequest);
}
