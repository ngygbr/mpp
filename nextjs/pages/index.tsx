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
    useDisclosure,
    Tabs, TabList, TabPanels, Tab, TabPanel, IconButton
} from '@chakra-ui/react'
import useSWR from 'swr'
import axios from "axios";
import Cookies from "js-cookie";
import {ReactNode, useEffect, useState} from "react";
import jwtDecode, {JwtPayload} from "jwt-decode";
import Card from "../components/card";
import CreditCard from "../components/creditCard"
import ACH from "../components/ach";
import { MdRefresh } from 'react-icons/md'

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

    function isExpired(expD: number | undefined) {
        if (tokenExp) {
            if (new Date(1000 * tokenExp) < new Date(Date.now())) {
                return true
            } else {
                return false
            }
        }
    }

    const ExpTokenCard = ({ children }: { children: ReactNode }) => (
        <Flex
            direction={"row"}
            justifyContent={"space-between"}
            alignItems={"center"}
        >
            <Box
                px={'20px'}
                py={'10px'}
                rounded={'md'}
                bg={'brand.200'}
                color={'brand.rejected'}
                maxW={'300px'}
                overflow={'hidden'}
                whiteSpace={'nowrap'}
                textOverflow={'ellipsis'}
            >
                {children}
            </Box>
            <IconButton
                aria-label='refresh-token'
                icon={<MdRefresh />}
                bg={"brand.200"}
                color={"brand.100"}
                _hover={
                    {bg: "brand.hover"}
                }
                onClick={getToken}
            />
        </Flex>
    );

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
                        isExpired(tokenExp) ?
                            <ExpTokenCard>{token}</ExpTokenCard>
                            :
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

            <Container
                maxW='container.xl'
                py={"4rem"}
            >
                <Flex
                    px={"2rem"}
                    w={"full"}
                    h={"full"}
                    direction={"column"}
                    justifyContent={"flex-start"}
                    alignItems={"center"}
                >
                    <Tabs isFitted variant='unstyled' w={"full"}>
                        <TabList gap={8} marginBottom={"4rem"}>
                            <Tab fontWeight={"600"} color={"brand.100"} bg={"brand.200"} borderRadius={"md"} _selected={{ color: 'brand.100', bg: 'brand.hover' }}>Credit Card</Tab>
                            <Tab fontWeight={"600"} color={"brand.100"} bg={"brand.200"} borderRadius={"md"} _selected={{ color: 'brand.100', bg: 'brand.hover' }}>Ach</Tab>
                            <Tab fontWeight={"600"} color={"brand.100"} bg={"brand.200"} borderRadius={"md"} _selected={{ color: 'brand.100', bg: 'brand.hover' }}>Apple Pay</Tab>
                            <Tab fontWeight={"600"} color={"brand.100"} bg={"brand.200"} borderRadius={"md"} _selected={{ color: 'brand.100', bg: 'brand.hover' }}>Google Pay</Tab>
                        </TabList>
                        <TabPanels>
                            <TabPanel bg={"brand.300"} color={"brand.100"}>
                                <CreditCard token={token} />
                            </TabPanel>
                            <TabPanel bg={"brand.300"} color={"brand.100"}>
                                <ACH token={token} />
                            </TabPanel>
                            <TabPanel bg={"brand.300"} color={"brand.100"}>
                                <p>AP</p>
                            </TabPanel>
                            <TabPanel bg={"brand.300"} color={"brand.100"}>
                                <p>GP</p>
                            </TabPanel>
                        </TabPanels>
                    </Tabs>
                </Flex>

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
                                    isExpired(tokenExp) ?
                                        <Text>Token is expired</Text>
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
