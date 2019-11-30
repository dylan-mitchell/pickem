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
  state = { active: false };

  handleClick = () => {
    const { active } = this.state;
    this.setState({ active: !active });
  };
  render() {
    return (
      <Navbar
        color={"white"}
        fixed={"top"}
        active={this.state.active}
        transparent={true}
      >
        <Navbar.Brand>
          <Navbar.Item>
            <Link style={brandStyle} to={ROUTES.LANDING}>
              Pickem
            </Link>
          </Navbar.Item>
          <Navbar.Burger
            active={this.state.active}
            onClick={this.handleClick}
          />
        </Navbar.Brand>
        <Navbar.Menu>
          <Navbar.Container>
            <Navbar.Item>
              <Link style={linkStyle} to={ROUTES.HOME}>
                Today
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

class NavigationNonAuth extends Component {
  state = { active: false };

  handleClick = () => {
    const { active } = this.state;
    this.setState({ active: !active });
  };

  render() {
    return (
      <Navbar
        color={"white"}
        fixed={"top"}
        active={this.state.active}
        transparent={false}
      >
        <Navbar.Brand>
          <Navbar.Item>
            <Link style={brandStyle} to={ROUTES.LANDING}>
              Pickem
            </Link>
          </Navbar.Item>
          <Navbar.Burger
            active={this.state.active}
            onClick={this.handleClick}
          />
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
  }
}

const linkStyle = {
  textDecoration: "none",
  color: "black"
};

const brandStyle = {
  fontFamily: "Permanent Marker",
  textDecoration: "none",
  color: "black",
  fontSize: "25px"
};

export default Navigation;
