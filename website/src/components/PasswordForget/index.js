import React, { Component } from "react";
import { Link } from "react-router-dom";
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
import Icon from "react-bulma-components/lib/components/icon";
import Level from "react-bulma-components/lib/components/level";

const PasswordForgetPage = () => (
  <Section>
    <Container>
      <Columns centered>
        <Columns.Column size="half">
          <Heading size={1} style={centerStyle}>
            Forgot Password?
          </Heading>
          <PasswordForgetForm />
        </Columns.Column>
      </Columns>
    </Container>
  </Section>
);
const INITIAL_STATE = {
  email: "",
  error: null,
  hidden: true
};

class PasswordForgetFormBase extends Component {
  constructor(props) {
    super(props);
    this.state = { ...INITIAL_STATE };
  }
  onSubmit = event => {
    const { email } = this.state;
    const hidden = false;
    this.props.firebase
      .doPasswordReset(email)
      .then(() => {
        this.setState({ hidden });
      })
      .catch(error => {
        this.setState({ error });
      });
    event.preventDefault();
  };
  onChange = event => {
    this.setState({ [event.target.name]: event.target.value });
  };
  getFormStyle = () => {
    return {
      display: this.state.hidden ? "block" : "none"
    };
  };
  getSuccessStyle = () => {
    return {
      display: this.state.hidden ? "none" : "block"
    };
  };

  render() {
    const { email, error } = this.state;
    const isInvalid = email === "";
    return (
      <React.Fragment>
        <div style={this.getFormStyle()}>
          <Field>
            <Heading size={4}>Email</Heading>
            <Control>
              <Input
                onChange={this.onChange}
                name="email"
                type="email"
                placeholder="Email address"
                value={this.state.email}
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
            Reset My Password
          </Button>
        </div>
        <div style={this.getSuccessStyle()}>
          <Level renderAs="nav">
            <Level.Item style={centerStyle}>
              <div>
                <Heading>Check your Inbox</Heading>
                <Heading>
                  <Icon>
                    <span className={"fas fa-check-circle fa-3x fa-fw"} />
                  </Icon>
                </Heading>
              </div>
            </Level.Item>
          </Level>
        </div>
      </React.Fragment>
    );
  }
}

const helpStyle = {
  color: "red"
};

const centerStyle = {
  textAlign: "center"
};

const PasswordForgetLink = () => (
  <p>
    <Link style={linkStyle} to={ROUTES.PASSWORD_FORGET}>
      Forgot Password?
    </Link>
  </p>
);

const linkStyle = {
  color: "black"
};

export default PasswordForgetPage;
const PasswordForgetForm = withFirebase(PasswordForgetFormBase);
export { PasswordForgetForm, PasswordForgetLink };
