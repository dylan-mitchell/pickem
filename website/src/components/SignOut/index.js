import React from "react";
import { withFirebase } from "../Firebase";

const SignOutButton = ({ firebase }) => (
  <a href="/SignIn" onClick={firebase.doSignOut} style={noStyle}>
    Sign Out
  </a>
);

const noStyle = {
  textDecoration: "none",
  color: "black"
};

export default withFirebase(SignOutButton);
