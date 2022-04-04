import type { NextPage } from 'next'
import {
    Button, Container
} from '@chakra-ui/react'
import {useState} from "react";

const Home: NextPage = () => {

    const [token, setToken] = useState()

    const login = async () => {
        const response = await fetch( "http://localhost:8080/login")
        const data = await response.json()
        setToken(data.token)
    }

  return (
      <Container
        size={'lg'}
      >
        <h1>Valami</h1>
      </Container>
  )
}

export default Home
