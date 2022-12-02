import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  env: NodeJS.ProcessEnv
}

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  res.status(200).json({ env: process.env })
}
