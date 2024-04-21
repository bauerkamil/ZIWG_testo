import { EnvironmentProfiles } from "../enums/EnviromentProfiles";

export const TOKEN_LINK = import.meta.env.VITE_TOKEN_LINK;
export const AUTH_LINK = import.meta.env.VITE_AUTH_LINK;
export const isTestEnv = import.meta.env.VITE_APP_ENV === EnvironmentProfiles.Test;