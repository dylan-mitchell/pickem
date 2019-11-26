import React from "react";
import { AuthUserContext, withAuthorization } from "../Session";
import UpdateUsernameForm from "../UpdateUsername";
import UpdateProfilePictureForm from "../UpdateProfilePicture";
import ProfilePicture from "../ProfilePicture";
import Section from "react-bulma-components/lib/components/section";
import Heading from "react-bulma-components/lib/components/heading";
import Container from "react-bulma-components/lib/components/container";
import Columns from "react-bulma-components/lib/components/columns";

const AccountPage = () => (
  <AuthUserContext.Consumer>
    {authUser => (
      <Section>
        <Container>
          <Columns centered>
            <Columns.Column size="half">
              <Heading size={1} style={centerStyle}>
                {authUser.displayName}
              </Heading>
              <div style={centerPic}>
                <ProfilePicture
                  imageURL={authUser.photoURL}
                  username={authUser.displayName}
                  size={128}
                />
              </div>
              <UpdateUsernameForm />
              <UpdateProfilePictureForm />
            </Columns.Column>
          </Columns>
        </Container>
      </Section>
    )}
  </AuthUserContext.Consumer>
);

const centerStyle = {
  textAlign: "center"
};

const centerPic = {
  marginLeft: "32%"
};

const condition = authUser => !!authUser;

export default withAuthorization(condition)(AccountPage);
