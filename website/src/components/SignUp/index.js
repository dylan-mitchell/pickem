import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";
import { compose } from "recompose";

import { withFirebase } from "../Firebase";
import * as ROUTES from "../../constants/routes";
import Section from "react-bulma-components/lib/components/section";
import Heading from "react-bulma-components/lib/components/heading";
import Container from "react-bulma-components/lib/components/container";
import {
  Field,
  Control,
  Input
} from "react-bulma-components/lib/components/form";
import Button from "react-bulma-components/lib/components/button";
import Columns from "react-bulma-components/lib/components/columns";

const SignUpPage = () => (
  <Section>
    <Container>
      <Columns centered>
        <Columns.Column size="half">
          <Heading size={1} style={centerStyle}>
            Sign Up
          </Heading>
          <SignUpForm />
        </Columns.Column>
      </Columns>
    </Container>
  </Section>
);

const INITIAL_STATE = {
  username: "",
  email: "",
  passwordOne: "",
  passwordTwo: "",
  error: null
};

class SignUpFormBase extends Component {
  constructor(props) {
    super(props);

    this.state = { ...INITIAL_STATE };
  }

  onSubmit = event => {
    const { username, email, passwordOne } = this.state;
    this.props.firebase
      .doCreateUserWithEmailAndPassword(email, passwordOne)
      .then(authUser => {
        this.props.firebase.doUpdateUsername(username);
        this.setState({ ...INITIAL_STATE });
        this.props.history.push(ROUTES.HOME);
      })
      .catch(error => {
        this.setState({ error });
      });
    event.preventDefault();
  };

  onChange = event => {
    this.setState({ [event.target.name]: event.target.value });
  };

  render() {
    const { username, email, passwordOne, passwordTwo, error } = this.state;

    const isInvalid =
      passwordOne !== passwordTwo ||
      passwordOne === "" ||
      email === "" ||
      username === "";
    return (
      <div>
        <Field>
          <Heading size={4}>Username</Heading>
          <Control>
            <Input
              onChange={this.onChange}
              name="username"
              type="username"
              placeholder="Username"
              value={username}
            />
          </Control>
        </Field>
        <Field>
          <Heading size={4}>Email</Heading>
          <Control>
            <Input
              onChange={this.onChange}
              name="email"
              type="email"
              placeholder="Email input"
              value={email}
            />
          </Control>
        </Field>
        <Field>
          <Heading size={4}>Password</Heading>
          <Control>
            <Input
              onChange={this.onChange}
              name="passwordOne"
              type="password"
              placeholder="Password"
              value={passwordOne}
            />
          </Control>
        </Field>
        <Field>
          <Heading size={4}>Confirm Password</Heading>
          <Control>
            <Input
              onChange={this.onChange}
              name="passwordTwo"
              type="password"
              placeholder="Password"
              value={passwordTwo}
            />
          </Control>
        </Field>
        {error && <p style={helpStyle}>{error.message}</p>}
        <Button
          disabled={isInvalid}
          type="submit"
          onClick={this.onSubmit}
          color="success"
          size="large"
        >
          Sign Up
        </Button>
      </div>
    );
  }
}

const helpStyle = {
  color: "red"
};

const centerStyle = {
  textAlign: "center"
};

const SignUpLink = () => (
  <Button.Group>
    <p>
      Don't have an account?{" "}
      <Link style={linkStyle} to={ROUTES.SIGN_UP}>
        Sign Up
      </Link>
    </p>
  </Button.Group>
);

const linkStyle = {
  color: "black"
};

const SignUpForm = compose(
  withRouter,
  withFirebase
)(SignUpFormBase);

export default SignUpPage;
export { SignUpForm, SignUpLink };
