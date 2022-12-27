

import React from 'react';
import axios from 'axios';
import { GetConfig } from '../../config';

type MyProps = {};
type MyState = { randomNumber: string };

const config = GetConfig();
class ButtonAPI extends React.Component<MyProps, MyState> {
  constructor(props: MyProps) {
    super(props);
    this.state = { randomNumber: "0" };
  }

  handleFetchError = () => {

  }
  getNumber = () => {
    axios.get(`${config.API_PROTOCOOL}://${config.API_EXTERNAL_URL}:${config.API_PORT}/api/v1/rand`, {
      headers: {
        "Accepts": "application/json",
      }
    }).then((response) => {
      const randomNumber = response.data.randomNumber;
      console.log(randomNumber);
      this.setState({ randomNumber: randomNumber })
    }).catch(error => {
      console.log(error)
      this.setState({ randomNumber: "NULL" })
      alert(`API at: ${config.API_PROTOCOOL}://${config.API_EXTERNAL_URL}:${config.API_PORT} is not online`)
    });
  };


  render() {
    const { randomNumber } = this.state;
    return (
      <div className="Button">
        <button onClick={this.getNumber} className="buttonAPI">RandomNumberAPI</button>
        <h1>{randomNumber}</h1>
      </div>
    );
  };
}

export default ButtonAPI;
