import ENV_LOCAL from "./env.local.json";
import ENV_K8S from "./env.k8s.json";

const isServer = typeof window === "undefined";

const EnvConfig = {
    local: ENV_LOCAL,
    k8s: ENV_K8S
};

export type EnvName = keyof typeof EnvConfig;

export const GetEnv = () => {
    let env = process.env.ENV_NAME;

    if (!isServer) {
        const element = document?.querySelector(
            'meta[name="env-name"]'
        ) as HTMLMetaElement;
        env = element?.content;
    }
    return env as EnvName | undefined;
}

export const GetConfig = () => {
    // console.log("scooby doo");
    const env = GetEnv();
    const config = {
        ...EnvConfig[env || "local"]
    };
    return config;
}
