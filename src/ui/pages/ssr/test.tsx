import { GetServerSideProps } from 'next'
import type { NextPage } from 'next'
import axios from 'axios';
import { GetConfig } from '../../config';

type PageProps = {
  randomNumber?: string
}

const config = GetConfig();

export const getServerSideProps: GetServerSideProps = async (context) => {

  const randomNumberReq = await axios.get(`${config.API_PROTOCOOL}://${config.API_INTERNAL_URL}:${config.API_PORT}/api/v1/rand`, {
    headers: {
      "Accepts": "application/json",
    }
  }).then((response) => {
    return response
  });
  const randomNumber = randomNumberReq.data.randomNumber

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
