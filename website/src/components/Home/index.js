import React from "react";
import { AuthUserContext, withAuthorization } from "../Session";
import ProfilePicture from "../ProfilePicture";

class HomePage extends React.Component {
  getAuthUser = a => {
    console.log(a);
  };
  render() {
    return (
      <AuthUserContext.Consumer>
        {authUser => (
          <div>
            <h1>Home Page</h1>
            <p>The Home Page is accessible by every signed in user.</p>
            <p>{this.getAuthUser(authUser)}</p>
            <ProfilePicture
              imageURL={authUser.photoURL}
              username={authUser.displayName}
              size={128}
            />
          </div>
        )}
      </AuthUserContext.Consumer>
    );
  }
}
const condition = authUser => !!authUser;
export default withAuthorization(condition)(HomePage);
