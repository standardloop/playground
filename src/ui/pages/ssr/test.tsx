import { GetServerSideProps } from 'next'
import type { NextPage } from 'next'
import axios from 'axios';

type PageProps = {
  randomNumber?: string
}

export const getServerSideProps: GetServerSideProps = async (context) => {

  const randomNumberReq = await axios.get(`${process.env.API_PROTOCOOL}://${process.env.API_INTERNAL_URL}:${process.env.API_PORT}/api/v1/rand`, {
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
