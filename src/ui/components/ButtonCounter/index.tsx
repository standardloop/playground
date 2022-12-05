import React from 'react';
import '../../styles/ButtonCounter.module.css'

type MyProps = {};
type MyState = { count: number };


class ButtonCounter extends React.Component<MyProps, MyState> {
  constructor(props: MyProps) {
    super(props);
    this.state = { count: 0 };
  }

  increment = () => {
    this.setState( {count: this.state.count + 1} )
  }

  render() {
    const { count } = this.state;
    return (
      <div className="buttonCounter">
        <button onClick={this.increment} className="buttonCounter">Counter</button>
        <h1>{ count }</h1>
      </div>
    );
  };
}

export default ButtonCounter;
