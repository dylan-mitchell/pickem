import React from "react";
import { AuthUserContext, withAuthorization } from "../Session";
import Matchup from "../Matchup";
import GroupRecord from "../GroupRecord";

const matchup = {
  option1: "Hawks",
  option1Image: "https://www.nba.com/assets/logos/teams/secondary/web/ATL.svg",
  option2: "Mavs",
  option2Image: "https://www.nba.com/assets/logos/teams/primary/web/DAL.svg",
  description: "The Mavs are favored by 3.5 points"
};

const groupInfo = {
  users: [
    {
      name: "Dylan",
      profilePic: "https://api.adorable.io/avatars/285/Dylan.png",
      wins: 5,
      losses: 3,
      choice: 2
    },
    {
      name: "Seth",
      profilePic: "https://api.adorable.io/avatars/285/Seth.png",
      wins: 3,
      losses: 5,
      choice: 0
    },
    {
      name: "Beau",
      profilePic: "https://api.adorable.io/avatars/285/Beau.png",
      wins: 2,
      losses: 6,
      choice: 1
    }
  ]
};

class TodayPage extends React.Component {
  render() {
    return (
      <React.Fragment>
        <AuthUserContext.Consumer>
          {authUser => <Matchup matchup={matchup} />}
        </AuthUserContext.Consumer>
        <GroupRecord groupInfo={groupInfo} matchup={matchup} />
      </React.Fragment>
    );
  }
}
const condition = authUser => !!authUser;
export default withAuthorization(condition)(TodayPage);
