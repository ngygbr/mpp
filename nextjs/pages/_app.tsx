import type { AppProps } from 'next/app'
import { ChakraProvider } from '@chakra-ui/react'
import { extendTheme } from "@chakra-ui/react"

const theme = extendTheme({
    colors: {
        brand: {
            100: "#EDF2F7",
            200: "#4A5568",
            300: "#2D3748",
            400: "#81E6D9",
            "settled": "#48BB78",
            "pending": "#ECC94B",
            "rejected": "#F56565",
            "hover": "#1A202C"
        },
    },
    styles: {
        global: () => ({
            body: {
                bg: "#2D3748",
            }
        })
    },
})

function MyApp({ Component, pageProps }: AppProps) {
  return (
      <ChakraProvider theme={theme}>
          <Component {...pageProps} />
      </ChakraProvider>
  )
}

export default MyApp
