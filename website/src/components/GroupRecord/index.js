import React from "react";
import Heading from "react-bulma-components/lib/components/heading";
import Level from "react-bulma-components/lib/components/level";
import Box from "react-bulma-components/lib/components/box";
import ProfilePicture from "../ProfilePicture";

class GroupRecord extends React.Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  getMatchupChoice = user => {
    if (user.choice === 1) {
      return (
        <ProfilePicture imageURL={this.props.matchup.option1Image} size={64} />
      );
    } else if (user.choice === 2) {
      return (
        <ProfilePicture imageURL={this.props.matchup.option2Image} size={64} />
      );
    } else {
      return <div style={{ height: "100px" }}></div>;
    }
  };

  getLevelItems = () => {
    return this.props.groupInfo.users.map(user => (
      <Level.Item style={{ textAlign: "center" }}>
        <Box>
          <Heading size={5}>{user.name}</Heading>
          <ProfilePicture
            imageURL={user.profilePic}
            username={user.name}
            size={64}
          />
          <Heading size={6}>
            {user.wins}-{user.losses}
          </Heading>
          {this.getMatchupChoice(user)}
        </Box>
      </Level.Item>
    ));
  };

  render() {
    return (
      <React.Fragment>
        <Level style={{ width: "100%" }} renderAs="nav">
          {this.getLevelItems()}
        </Level>
      </React.Fragment>
    );
  }
}

export default GroupRecord;
