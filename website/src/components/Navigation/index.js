import React, { Component } from "react";
import * as ROUTES from "../../constants/routes";
import SignOutButton from "../SignOut";
import Navbar from "react-bulma-components/lib/components/navbar";
import { Link } from "react-router-dom";

import { AuthUserContext } from "../Session";
const Navigation = () => (
  <div>
    <AuthUserContext.Consumer>
      {authUser =>
        authUser ? <NavigationAuth user={authUser} /> : <NavigationNonAuth />
      }
    </AuthUserContext.Consumer>
  </div>
);

class NavigationAuth extends Component {
  render() {
    return (
      <Navbar
        color={"primary"}
        fixed={"top"}
        active={false}
        transparent={false}
      >
        <Navbar.Brand>
          <Navbar.Item>
            <Link to={ROUTES.LANDING}>
              <img
                src="https://bulma.io/images/bulma-logo.png"
                alt="Bulma: a modern CSS framework based on Flexbox"
                width="112"
                height="28"
              />
            </Link>
          </Navbar.Item>
          <Navbar.Burger />
        </Navbar.Brand>
        <Navbar.Menu>
          <Navbar.Container>
            <Navbar.Item>
              <Link style={linkStyle} to={ROUTES.HOME}>
                Home
              </Link>
            </Navbar.Item>
            <Navbar.Item>
              <Link style={linkStyle} to={ROUTES.ACCOUNT}>
                Account
              </Link>
            </Navbar.Item>
          </Navbar.Container>
          <Navbar.Container position="end">
            <Navbar.Item>
              <SignOutButton />
            </Navbar.Item>
          </Navbar.Container>
        </Navbar.Menu>
      </Navbar>
    );
  }
}

const NavigationNonAuth = () => (
  <Navbar color={"primary"} fixed={"top"} active={false} transparent={false}>
    <Navbar.Brand>
      <Navbar.Item>
        <Link to={ROUTES.LANDING}>
          <img
            src="https://bulma.io/images/bulma-logo.png"
            alt="Bulma: a modern CSS framework based on Flexbox"
            width="112"
            height="28"
          />
        </Link>
      </Navbar.Item>
      <Navbar.Burger />
    </Navbar.Brand>
    <Navbar.Menu>
      <Navbar.Container position="end">
        <Navbar.Item>
          <Link style={linkStyle} to={ROUTES.SIGN_IN}>
            Sign In
          </Link>
        </Navbar.Item>
      </Navbar.Container>
    </Navbar.Menu>
  </Navbar>
);

const linkStyle = {
  textDecoration: "none",
  color: "black"
};

export default Navigation;
