import React from "react"
import { Table, Button } from 'react-bootstrap';
import { Transaction } from "../model/types";

import TranactionFormContent from "../contents/transactionForm";


export default function TransactionFormContainer() { 
  return TranactionFormContent({})
}
