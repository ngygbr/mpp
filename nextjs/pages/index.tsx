import type { NextPage } from 'next'
import {
    Box,
    Button,
    Container,
    Drawer,
    DrawerBody,
    DrawerContent,
    DrawerHeader,
    DrawerOverlay,
    Flex,
    Grid,
    GridItem,
    Spinner,
    Text,
    useDisclosure
} from '@chakra-ui/react'
import useSWR from 'swr'
import axios from "axios";
import Cookies from "js-cookie";
import {ReactNode, useEffect, useState} from "react";
import jwtDecode, {JwtPayload} from "jwt-decode";
import Card from "../components/card";

const Home: NextPage = () => {
    const [token, setToken] = useState("")
    const [tokenExp, setTokenExp] = useState<undefined | number>()
    const { isOpen, onOpen, onClose } = useDisclosure()

    const fetcher = (token: string) => axios.get("http://localhost:8080/api/transactions", {
        method: "GET",
        headers: {
            "Authorization": token
        }
    }).then(res => res.data)

    const getToken = async () => {
        const response = await fetch( "http://localhost:8080/login")
        const data = await response.json()
        setToken(data.token)
        Cookies.set("mpp.AuthToken", data.token)
        const decoded = jwtDecode<JwtPayload>(data.token);
        setTokenExp(decoded.exp)
    }

    useEffect(() => {
        if (tokenExp) {
            const expDate = new Date(1000 * tokenExp)
            const now = new Date(Date.now())

            if (expDate < now) {
                setToken("");
                Cookies.remove('mpp.AuthToken')
            }
        }
    }, [token]);

    const { data, error } = useSWR(token, token => fetcher(token))

    if (error) return <div>failed to load</div>

    return (
        <>
            <Box
                bg={'brand.300'}
            >
                <Flex
                    h={'60px'}
                    maxH={'60px'}
                    px={'20px'}
                    py={'10px'}
                    alignItems={'center'}
                    justifyContent={'space-between'}
                >
                    {token ?
                        <TokenCard>{token}</TokenCard>
                        :
                        <Button
                            onClick={getToken}
                            bg={"brand.200"}
                            color={"brand.100"}
                            _hover={
                                {bg: "brand.hover"}
                            }
                        >
                            Get Token
                        </Button>
                    }

                    <Button
                        bg={"brand.200"}
                        color={"brand.100"}
                        _hover={
                            {bg: "brand.hover"}
                        }
                        onClick={onOpen}
                    >
                        Transactions
                    </Button>
                </Flex>
            </Box>

            <Container maxW='container.xl'>
                TODO

            </Container>

            <Drawer
                placement="right"
                onClose={onClose}
                isOpen={isOpen}
                size={"xl"}
            >
                <DrawerOverlay/>
                <DrawerContent
                    bg={"brand.200"}
                    color={"brand.100"}
                >
                    <DrawerHeader
                        borderBottomWidth='1px'
                        borderColor={"brand.200"}
                    >
                        Transactions
                    </DrawerHeader>
                    <DrawerBody>
                        {data ?
                            <Grid templateColumns={"repeat(3, 1fr)"} gap={4}>
                                {data.transactions ?
                                    data.transactions.map((transaction: any) => (
                                        <GridItem w='100%' key={transaction.id}>
                                            <Card
                                                key={transaction.id}
                                                token={token}
                                                transaction={transaction}
                                            />
                                        </GridItem>
                                    ))
                                    :
                                    <Text>No transactions yet...</Text>
                                }
                            </Grid>
                            :
                            <Flex
                                marginTop={"-60px"}
                                justifyContent={"center"}
                                alignItems={"center"}
                                w={"full"}
                                h={"full"}
                            >
                                <Spinner size='lg' />
                            </Flex>
                        }
                    </DrawerBody>
                </DrawerContent>
            </Drawer>
        </>
    )
}

export default Home

const TokenCard = ({ children }: { children: ReactNode }) => (
    <Box
        px={'20px'}
        py={'10px'}
        rounded={'md'}
        bg={'brand.200'}
        color={'brand.100'}
        maxW={'300px'}
        overflow={'hidden'}
        whiteSpace={'nowrap'}
        textOverflow={'ellipsis'}
    >
        {children}
    </Box>
);

const dummyT = [
    {"id":"c95l5iu49b3jj2kuu990","status":"settled","payment_method_type":"creditcard","payment_method":{"credit_card":{"card_number":"544416******3444","holder_name":"John Lécci","exp_date":"05/25","cvc":"444"}},"amount":11111,"billing_address":{"first_name":"John","last_name":"Doe","postal_code":"1111","city":"Szeged","address_line_1":"Example street 69.","email":"example@github.com","phone":"5555555555"},"created_at":"2022-04-04T22:19:23.997539+02:00","updated_at":"2022-04-04T22:19:23.997539+02:00"}
]

function printValues(obj: { [x: string]: any; }, arr: Array<string>) {

    for (const key in obj) {
        if (typeof obj[key] === "object") {
            printValues(obj[key], arr);
        } else {
            arr.push(obj[key])
        }
    }

    console.log(arr)
    return arr
}
