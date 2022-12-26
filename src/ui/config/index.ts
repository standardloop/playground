import ENV_LOCAL from "./env.local.json";
import ENV_K8S from "./env.k8s.json";


const isServer = typeof window === "undefined";

const EnvConfig = {
    local: ENV_LOCAL,
    k8s: ENV_K8S
};

export type EnvName = keyof typeof EnvConfig;

export const getEnv = () => {
    let env = process.env.ENV_NAME;
    return env as EnvName | undefined;
}

export const getConfig = () => {
    const env = getEnv();
    const config = {
        ...EnvConfig[env || "local"]
    };
}
