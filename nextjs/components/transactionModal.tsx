import {
    Accordion, AccordionButton, AccordionIcon, AccordionItem, AccordionPanel,
    Box,
    Button,
    HStack,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent, ModalFooter,
    ModalHeader,
    ModalOverlay,
    Stack, Tag,
    Text
} from "@chakra-ui/react";
import React, {PropsWithChildren} from "react";
import axios from "axios";
import {randomUUID} from "crypto";
import useSWR from "swr";

type Props = {
    token: string,
    isOpened: boolean;
    closeModal: () => void;
    transaction: {
        id: string,
        status: string,
        payment_method_type: string,
        payment_method: object,
        amount: number,
        billing_address: {
            first_name: string,
            last_name: string,
            postal_code: string,
            city: string,
            address_line_1: string,
            email: string,
            phone: string
        }
        created_at: string,
        updated_at: string,
    };
} & Record<string, any>;

const TransactionModal = ( {token, isOpened,closeModal,transaction}: PropsWithChildren<Props>) => {

    const settleTransaction = () => axios.get('http://localhost:8080/api/transaction/'+ transaction.id + '/settle', {
        headers: {
            "Authorization": token
        }
    })

    const rejectTransaction = () => axios.get('http://localhost:8080/api/transaction/'+ transaction.id + '/reject', {
        headers: {
            "Authorization": token
        }
    })

    const deleteTransaction = () => axios.delete('http://localhost:8080/api/transaction/'+ transaction.id, {
        headers: {
            "Authorization": token
        }
    })

    return (
        <Modal
            isOpen={isOpened}
            onClose={closeModal}
        >
            <ModalOverlay/>
            <ModalContent
                bg={"brand.300"}
                color={"brand.100"}
            >
                <ModalHeader
                    borderBottomWidth='1px'
                    borderColor={"brand.200"}
                >Transaction</ModalHeader>
                <ModalCloseButton/>

                <ModalBody
                    py={"1.5rem"}
                >

                    <Stack>
                        <Text><Tag bg={"brand.200"} color={"brand.100"} marginRight={"10px"}>ID</Tag>{transaction.id}</Text>
                        <Text><Tag bg={"brand.200"} color={"brand.100"}
                                   marginRight={"10px"}>Status</Tag>{statusFormat(transaction.status)}</Text>
                        <Text><Tag bg={"brand.200"} color={"brand.100"}
                                   marginRight={"10px"}>Type</Tag>{typeFormat(transaction.payment_method_type)}</Text>
                        <Text><Tag bg={"brand.200"} color={"brand.100"}
                                   marginRight={"10px"}>Amount</Tag>{transaction.amount}</Text>

                        <Accordion allowToggle
                                   w={"full"}
                                   bg={"brand.200"}
                                   color={"brand.100"}
                                   borderRadius={"md"}
                        >
                            <AccordionItem
                                border={"none"}
                            >
                                <h2>
                                    <AccordionButton>
                                        <Box
                                            flex='1'
                                            textAlign='left'
                                        >
                                            Payment Data
                                        </Box>
                                        <AccordionIcon/>
                                    </AccordionButton>
                                </h2>
                                <AccordionPanel>
                                    <Stack>
                                        {printValues(transaction.payment_method, []).map((item) =>
                                            <Text key={item} fontSize={"14px"}>{item}</Text>)}
                                    </Stack>
                                </AccordionPanel>
                            </AccordionItem>
                            <AccordionItem
                                border={"none"}
                            >
                                <h2>
                                    <AccordionButton>
                                        <Box
                                            flex='1'
                                            textAlign='left'
                                        >
                                            Address
                                        </Box>
                                        <AccordionIcon/>
                                    </AccordionButton>
                                </h2>
                                <AccordionPanel>
                                    <Stack>
                                        <Text><Tag bg={"brand.300"} color={"brand.100"} marginRight={"10px"}>First
                                            Name</Tag>{transaction.billing_address.first_name}</Text>
                                        <Text><Tag bg={"brand.300"} color={"brand.100"} marginRight={"10px"}>Last
                                            Name</Tag>{transaction.billing_address.last_name}</Text>
                                        <Text><Tag bg={"brand.300"} color={"brand.100"} marginRight={"10px"}>Postal
                                            Code</Tag>{transaction.billing_address.postal_code}</Text>
                                        <Text><Tag bg={"brand.300"} color={"brand.100"}
                                                   marginRight={"10px"}>City</Tag>{transaction.billing_address.city}</Text>
                                        <Text><Tag bg={"brand.300"} color={"brand.100"} marginRight={"10px"}>Address
                                            Line</Tag>{transaction.billing_address.address_line_1}</Text>
                                        <Text><Tag bg={"brand.300"} color={"brand.100"}
                                                   marginRight={"10px"}>Email</Tag>{transaction.billing_address.email}
                                        </Text>
                                        <Text><Tag bg={"brand.300"} color={"brand.100"}
                                                   marginRight={"10px"}>Phone</Tag>{transaction.billing_address.phone}
                                        </Text>
                                    </Stack>
                                </AccordionPanel>
                            </AccordionItem>
                        </Accordion>

                        <Text><Tag bg={"brand.200"} color={"brand.100"}
                                   marginRight={"10px"}>Created</Tag>{readableDate(transaction.created_at)}</Text>
                        <Text><Tag bg={"brand.200"} color={"brand.100"} marginRight={"10px"}>Last
                            Updated</Tag>{readableDate(transaction.updated_at)}</Text>
                    </Stack>

                </ModalBody>

                <ModalFooter
                    display={"flex"}
                    dir={"row"}
                    alignItems={"center"}
                    justifyContent={"space-between"}
                    borderTopWidth='1px'
                    borderColor={"brand.200"}
                >
                    <HStack>
                        <Button
                            bg={"brand.200"}
                            color={"brand.100"}
                            _hover={
                                {bg: "brand.hover"}
                            }
                            onClick={settleTransaction}
                        >Settle</Button>
                        <Button
                            bg={"brand.200"}
                            color={"brand.100"}
                            _hover={
                                {bg: "brand.hover"}
                            }
                            onClick={rejectTransaction}
                        >Reject</Button>
                    </HStack>
                    <Button
                        bg={"brand.200"}
                        color={"brand.100"}
                        _hover={
                            {bg: "brand.hover"}
                        }
                        onClick={deleteTransaction}
                    >Delete</Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )
};

export default TransactionModal;

function printValues(obj: { [x: string]: any; }, arr: Array<string>) {

    for (const key in obj) {
        if (typeof obj[key] === "object") {
            printValues(obj[key], arr);
        } else {
            arr.push(obj[key])
        }
    }

    return arr
}

function typeFormat(type: string) {
    let formattedType

    switch(type) {
        case "creditcard":
            formattedType = "Credit Card";
            break;
        case "ach":
            formattedType = "Ach";
            break;
        case "apple_pay":
            formattedType = "Apple Pay";
            break;
        case "google_pay":
            formattedType = "Google Pay";
            break;
    }

    return formattedType
}

function statusFormat(type: string) {
    let formattedStatus

    switch(type) {
        case "pending_settlement":
            formattedStatus = "Pending Settlement";
            break;
        case "rejected":
            formattedStatus = "Rejected";
            break;
        case "settled":
            formattedStatus = "Settled";
            break;
    }

    return formattedStatus
}

function readableDate(date: string) {
    const d = new Date(date);
    return d.toDateString()
}
