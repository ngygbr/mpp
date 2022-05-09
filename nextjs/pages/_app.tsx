import type { AppProps } from "next/app";
import {ChakraProvider, useMediaQuery} from "@chakra-ui/react";
import { extendTheme } from "@chakra-ui/react";
import Small from "../components/small";

const theme = extendTheme({
  colors: {
    brand: {
      100: "#EDF2F7",
      200: "#4A5568",
      300: "#2D3748",
      400: "#81E6D9",
      settled: "#48BB78",
      pending: "#ECC94B",
      rejected: "#F56565",
      hover: "#1A202C",
    },
  },
  styles: {
    global: () => ({
      body: {
        bg: "#2D3748",
      },
    }),
  },
});

function MyApp({ Component, pageProps }: AppProps) {
  const [isLargerThan1080] = useMediaQuery('(min-width: 380px)')

  return (
      <ChakraProvider theme={theme}>
        {isLargerThan1080 ?
              <Component {...pageProps} />
            :
              <Small/>
        }
      </ChakraProvider>
  );
}

export default MyApp;
