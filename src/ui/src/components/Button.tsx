import React from 'react';
import './Button.css';
import axios from 'axios';

type MyProps = {};
type MyState = {randomNumber: string, count: number};

class Button extends React.Component<MyProps, MyState> {
  constructor(props: MyProps) {
    super(props);
    this.state = { randomNumber: "0", count: 0};
  }

  handleFetchError = () => {

  }
  getNumber = () => {
    axios.get("/api/v1/rand/", {
    headers: {
      "accepts": "application/json"
    }}).then((response) => {
        const randomNumber = response.data.randomNumber;
        console.log(randomNumber);
        this.setState({randomNumber: randomNumber})
      }).catch(error => {
          console.log(error)
          this.setState({randomNumber: "API is not online"})
          alert("API is not online")
      });
  };
  increment = () => {
    this.setState( {count: this.state.count + 1} )
  }

  render() {
    const { randomNumber } = this.state;
    const { count } = this.state;
    return (
      <div className="Button">
        <h1>playground</h1>
        <button onClick={this.increment} className="testButton">Counter</button>
        <h1>{ count }</h1>
        <button onClick={this.getNumber} className="testButton">RandomNumberAPI</button>
        <h1>{ randomNumber }</h1>
      </div>
    );
  };
}

export default Button;
