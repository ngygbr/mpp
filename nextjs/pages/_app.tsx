import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { ChakraProvider } from '@chakra-ui/react'
import Nav from "../components/navbar"

function MyApp({ Component, pageProps }: AppProps) {
  return (
      <ChakraProvider>
          <Nav />
          <Component {...pageProps} />
      </ChakraProvider>
  )
}

export default MyApp
