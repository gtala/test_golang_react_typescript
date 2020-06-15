import React from 'react';
import logo from './logo.svg';
import './App.css';
import TransactionContainer from './containers/transactionsListContainer'
import TransactionFormContainer from './containers/transactionFormContainer'

//@ts-ignore
import {BrowserRouter as Router,Switch,Route,Link} from "react-router-dom";

export default function App() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/about">About</Link>
            </li>
            <li>
              <Link to="/transactions">Transactions</Link>
            </li>
          </ul>
        </nav>

        {/* A <Switch> looks through its children <Route>s and
            renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/about">
            <About />
          </Route>
          <Route
            path="/transactions"
            component={() => (
              //@ts-ignore
              <TransactionContainer />
            )}
          />
          <Route
            path="/form"
            component={() => (
              //@ts-ignore
              <TransactionFormContainer />
            )}
          />
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

function Home() {
  return <h2>Home</h2>;
}

function About() {
  return <h2>About</h2>;
}

