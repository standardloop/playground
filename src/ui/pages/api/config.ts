import type { NextApiRequest, NextApiResponse } from 'next'
import type { EnvName } from "../../config"
import { getEnv } from '../../config'

type Data = {
    env: EnvName | undefined
}

export default function handler(
    req: NextApiRequest,
    res: NextApiResponse<Data>
) {
    res.status(200).json({ env: getEnv() })
}
