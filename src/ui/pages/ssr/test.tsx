import { GetServerSideProps } from 'next'
import type { NextPage } from 'next'
import axios from 'axios';
import { GetConfig } from '../../config';

type PageProps = {
  randomNumber?: string
}

const config = GetConfig();

export const getServerSideProps: GetServerSideProps = async (context) => {
  const url = `${config.API_PROTOCOOL}://${config.API_INTERNAL_URL}:${config.API_PORT}/api/v1/rand`
  console.log(url);
  const randomNumberReq = await axios.get(url, {
    headers: {
      "Accepts": "application/json",
    }
  }).then((response) => {
    return response
  }).catch((err) => {
    return { "data": { "randomNumber": err.code } }
  });
  //console.log(randomNumberReq);
  const randomNumber = randomNumberReq.data.randomNumber
  context.res.setHeader(
    'Cache-Control',
    'public, s-maxage=10, stale-while-revalidate=59'
  )
  const _props: PageProps = {
    randomNumber: randomNumber
  }
  return { props: _props }
};

const Test: NextPage<PageProps> = (props) => {
  return (
    <div className="foo">
      <h1>{props.randomNumber}</h1>
    </div>
  );
};

export default Test
