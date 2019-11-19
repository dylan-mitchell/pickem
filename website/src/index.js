import React from 'react';
import ReactDOM from 'react-dom';
import {
	// Button,
	Navbar,
} from 'react-bulma-components/lib';

function NavigationBar(props) {
  return (
		<Navbar
        color='primary'
        fixed='Fixed'
        active='false'
        transparent='false'
      >
        <Navbar.Brand>
          <Navbar.Item renderAs="a" href="#">
            <img src="../logo512.png" alt="Logo"/>
          </Navbar.Item>
          <Navbar.Burger />
        </Navbar.Brand>
        <Navbar.Menu >
          <Navbar.Container>
            <Navbar.Item dropdown hoverable href="#">
              <Navbar.Link false>
                First
              </Navbar.Link>
              <Navbar.Dropdown>
                <Navbar.Item href="#">
                  Subitem 1
                </Navbar.Item>
                <Navbar.Item href="#">
                  Subitem 2
                </Navbar.Item>
              </Navbar.Dropdown>
            </Navbar.Item>
            <Navbar.Item href="#">
              Second
            </Navbar.Item>
          </Navbar.Container>
          <Navbar.Container position="end">
            <Navbar.Item href="#">
                  At the end
            </Navbar.Item>
          </Navbar.Container>
        </Navbar.Menu>
      </Navbar>
  );
}

class App extends React.Component {
	constructor(props) {
    super(props);
    this.state = {
    };
  }

  render() {
    return (
			<NavigationBar></NavigationBar>
    );
  }
}

// ========================================

ReactDOM.render(
  <App />,
  document.getElementById('root')
);
