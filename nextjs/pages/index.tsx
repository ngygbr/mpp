import type { NextPage } from 'next'
import {
    Box,
    Button,
    Container,
    Drawer,
    Stack,
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
    Tabs, TabList, TabPanels, Tab, TabPanel, IconButton, HStack, useToast
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
import {any} from "prop-types";
import TokenPayment from "../components/tokenPayment";

const Home: NextPage = () => {
    const [token, setToken] = useState("")
    const [tokenExp, setTokenExp] = useState<undefined | number>()
    const { isOpen, onOpen, onClose } = useDisclosure()
    const { data: transactions, mutate, error } = useSWR(token, token => fetcher(token))
    const toast = useToast()

    const fetcher = (token: string) => axios.get("http://localhost:8080/api/transactions", {
        method: "GET",
        headers: {
            "Authorization": token
        }
    }).then((res) => res.data)

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

    const deleteAllTransaction = async () => {
        try {
            const resp = await axios.delete('http://localhost:8080/api/transactions', {
                headers: {
                    "Authorization": token
                }
            })
            if (resp.status == 200) {
                toast({
                    title: resp.data.message,
                    status: "success",
                    isClosable: true,
                })
            } else {
                toast({
                    title: resp.data.message,
                    status: "error",
                    isClosable: true,
                })
            }
            await mutate(token)
        } catch (error: any) {
            if (error.response){
                toast({
                    title: error.response.data.message,
                    status: "error",
                    isClosable: true,
                })
            } else {
                toast({
                    title: error.message,
                    status: "error",
                    isClosable: true,
                })
            }
        }
    }

    if (error) return <div>failed to load</div>

    if (!token) return <><Box
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
        </Flex>
    </Box>
        <Container
            maxW='container.xl'
            py={"4rem"}
            display={"flex"}
            justifyContent={"center"}
            bg={"brand.300"}
        >
            <Text
                fontWeight={"600"}
                color={"brand.100"}
            >Get Access Token First</Text>
        </Container>
    </>

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
                            <Button
                                onClick={getToken}
                                rightIcon={<MdRefresh />}
                                bg={"brand.200"}
                                color={"brand.rejected"}
                                _hover={
                                    {bg: "brand.hover"}
                                }
                            >
                                Refresh Token
                            </Button>
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
                { (!isExpired(tokenExp)) ?
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
                                <CreditCard token={token} mutator={mutate}/>
                            </TabPanel>
                            <TabPanel bg={"brand.300"} color={"brand.100"}>
                                <ACH token={token} mutator={mutate}/>
                            </TabPanel>
                            <TabPanel bg={"brand.300"} color={"brand.100"}>
                                <TokenPayment token={token} mutator={mutate} type={"applepay"}/>
                            </TabPanel>
                            <TabPanel bg={"brand.300"} color={"brand.100"}>
                                <TokenPayment token={token} mutator={mutate} type={"googlepay"}/>
                            </TabPanel>
                        </TabPanels>
                    </Tabs>
                </Flex>
                    :
                    <Flex
                        justifyContent={"center"}
                    >
                    <Text
                        fontWeight={"600"}
                        color={"brand.100"}
                    >Your Token Is Expired</Text>
                    </Flex>
                }

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
                        display={"flex"}
                        flexDirection={"row"}
                        justifyContent={"space-between"}
                        alignItems={"center"}
                    >
                        <Text>Transactions</Text>
                        {transactions ?
                            transactions.transactions ?
                                <Button
                                    bg={"brand.300"}
                                    color={"brand.100"}
                                    _hover={
                                        {bg: "brand.hover"}
                                    }
                                    onClick={deleteAllTransaction}
                                >
                                    Delete all
                                </Button>
                                :
                                <></>
                            :
                            <></>
                        }
                    </DrawerHeader>
                    <DrawerBody>
                        {transactions ?
                            <Grid templateColumns={"repeat(3, 1fr)"} gap={4}>
                                {transactions.transactions ?
                                    transactions.transactions.map((transaction: any) => (
                                        <GridItem w='100%' key={transaction.id}>
                                            <Card
                                                key={transaction.id}
                                                token={token}
                                                mutator={mutate}
                                                transaction={transaction}
                                            />
                                        </GridItem>
                                    ))
                                    :
                                    isExpired(tokenExp) ?
                                        <GridItem
                                            colSpan={3}
                                            justifySelf={"center"}
                                        > <Text>Token is expired</Text> </GridItem>
                                        :
                                        <GridItem
                                            colSpan={3}
                                            justifySelf={"center"}
                                        > <Text>No transactions yet...</Text> </GridItem>
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
