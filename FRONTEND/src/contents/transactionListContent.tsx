import React, { useState } from "react"
import { Table, Button , Modal, Form} from 'react-bootstrap';
import { Transaction } from "../model/types";
import { goToForm } from "../utils/historyRouter";
import { Link } from "react-router-dom";

export default function TranactionListContent(props: any, newtransaction: Transaction) {


  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);


  const render = () => (
    <>
      <Button
        variant="primary"
        onClick={() => {
          setShow(true);
        }}
      >
        New
      </Button>
      {""}
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>#Transaction ID</th>
            <th>Account origin alias</th>
            <th>Account origin user name</th>
            <th>Account destination alias</th>
            <th>Account destination user name</th>
            <th>Amount</th>
            <th>DateTime</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {props.transactions.map((t: Transaction) => (
            <tr>
              <td>{t.id}</td>
              <td>{t.source.alias}</td>
              <td>{t.source.owner.name}</td>
              <td>{t.destination.alias}</td>
              <td>{t.destination.owner.name}</td>
              <td>{t.amount}</td>
              <td>{t.dateTime}</td>
            </tr>
          ))}
        </tbody>
      </Table>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>New Transaction</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group controlId="amount">
              <Form.Label>Amount</Form.Label>
              {
                //@ts-ignore
              }
              <Form.Control type="text" placeholder="Enter Amount" value={newtransaction.amount}>
                
              </Form.Control>
            </Form.Group>
            <Form.Group controlId="origin account">
              <Form.Label>origin account</Form.Label>
              <Form.Control as="select">
                <option>1</option>
                <option>2</option>
              </Form.Control>
            </Form.Group>
            <Form.Group controlId="destination account">
              <Form.Label>destination account</Form.Label>
              <Form.Control as="select">
              <option>1</option>
                <option>2</option>
              </Form.Control>
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={() => { props.saveChange({})}}>
            Save Changes
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );

  
  return render()
}