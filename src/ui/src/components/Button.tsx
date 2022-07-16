import React from 'react';

import './Button.css';

type MyProps = {};
type MyState = {randomNumber: string};

class Button extends React.Component<MyProps, MyState> {
  constructor(props: MyProps) {
    super(props);
    this.state = { randomNumber: "0" };
  }
  
  getNumber = () => {
    fetch("/api/v1/rand/",)
      .then(response => response.json())
      .then(responseJson => {
        const randomNumber = responseJson.randomNumber;
        console.log(randomNumber);
        this.setState({randomNumber: randomNumber})
      })
  };

  render() {
    const { randomNumber } = this.state;
    return (
      <div className="Button">
        <h1>Whattup</h1>
        <button onClick={this.getNumber} className="testButton">CLICK ME!</button>
        <h1>{ randomNumber }</h1>
      </div>
    );
  };

}

export default Button;
