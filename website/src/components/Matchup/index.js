import React from "react";
import Heading from "react-bulma-components/lib/components/heading";
import Image from "react-bulma-components/lib/components/image";
import Columns from "react-bulma-components/lib/components/columns";
import Level from "react-bulma-components/lib/components/level";
import Button from "react-bulma-components/lib/components/button";
import Icon from "react-bulma-components/lib/components/icon";

class Matchup extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      selected: 0,
      lockedIn: 0
    };
  }

  handleClickOption1 = () => {
    if (this.state.lockedIn === 0) {
      const selected = 1;
      this.setState({ selected });
    }
  };

  handleClickOption2 = () => {
    if (this.state.lockedIn === 0) {
      const selected = 2;
      this.setState({ selected });
    }
  };

  lockItIn = () => {
    const lockedIn = this.state.selected;
    this.setState({ lockedIn });
  };

  unlockItIn = () => {
    const lockedIn = 0;
    this.setState({ lockedIn });
  };

  getStyle = option => {
    if (option === this.state.selected) {
      var border = "1px solid #ccc";
      var boxShadow = "5px 10px 18px #ccc";
    } else {
      border = "";
      boxShadow = "";
    }

    if (option === this.state.lockedIn) {
      border = "1px solid #FFD700";
      boxShadow = "5px 10px 18px #FFD700";
    }

    return {
      border: border,
      padding: "10px",
      borderRadius: "16px",
      boxShadow: boxShadow
    };
  };

  getButton = () => {
    if (this.state.lockedIn === 0) {
      return (
        <Button color={"black"} outlined size={"large"} onClick={this.lockItIn}>
          <Icon>
            <span className="fas fa-lock" />
          </Icon>
          <Icon style={{ display: "none" }}>
            <span className="fas fa-unlock" />
          </Icon>
          <span />
          Lock it In
        </Button>
      );
    } else {
      return (
        <Button color={"black"} size={"large"} onClick={this.unlockItIn}>
          <Icon style={{ display: "none" }}>
            <span className="fas fa-unlock" />
          </Icon>
          <Icon>
            <span className="fas fa-unlock" />
          </Icon>
          <span />
          Unlock
        </Button>
      );
    }
  };

  render() {
    return (
      <React.Fragment>
        <Heading style={{ textAlign: "center" }}>
          {this.props.matchup.option1} vs {this.props.matchup.option2}
        </Heading>
        <Heading size={5} style={{ textAlign: "center" }}>
          {this.props.matchup.description}
        </Heading>
        <Columns centered>
          <Columns.Column size="one-third">
            <Level>
              <Level.Item>
                <div style={this.getStyle(1)} onClick={this.handleClickOption1}>
                  <Image src={this.props.matchup.option1Image} size={128} />
                </div>
              </Level.Item>
            </Level>
          </Columns.Column>
          <Columns.Column size="one-third">
            <Level>
              <Level.Item>
                <div style={this.getStyle(2)} onClick={this.handleClickOption2}>
                  <Image src={this.props.matchup.option2Image} size={128} />
                </div>
              </Level.Item>
            </Level>
          </Columns.Column>
        </Columns>
        <Level>
          <Level.Item>{this.getButton()}</Level.Item>
        </Level>
      </React.Fragment>
    );
  }
}

export default Matchup;
