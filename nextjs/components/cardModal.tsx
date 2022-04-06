import {
    Accordion, AccordionButton, AccordionIcon, AccordionItem, AccordionPanel,
    Box,
    Button, FormControl, FormErrorMessage, FormLabel, Grid, GridItem,
    HStack, Input,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent, ModalFooter,
    ModalHeader,
    ModalOverlay,
    Stack, Tag,
    Text, useToast
} from "@chakra-ui/react";
import React, {PropsWithChildren, useState, useCallback} from "react";
import axios from "axios";
import {useForm} from "react-hook-form";
import useSWR, {KeyedMutator, useSWRConfig} from "swr";

type Props = {
    token: string,
    mutator:  KeyedMutator<any>,
    isOpened: boolean;
    closeModal: () => void;
} & Record<string, any>;

function base64ToHex(str: string) {
    const raw = atob(str);
    let result = '';
    for (let i = 0; i < raw.length; i++) {
        const hex = raw.charCodeAt(i).toString(16);
        result += (hex.length === 2 ? hex : '0' + hex);
    }
    return result.toUpperCase();
}

const CardModal = ( {token, mutator, isOpened,closeModal, encryptedCard, encryptedKey, onEncryptedCardChange, onEncryptedKeyChange}: PropsWithChildren<Props>) => {
    const toast = useToast()

    const addCard = async (values: any | undefined) => {
        try {
            const resp = await axios.post(
                'http://localhost:8080/encryptcard',
                {
                        "credit_card": {
                            "card_number": values.card_number,
                            "holder_name": values.holder,
                            "exp_date": values.exp_date,
                            "cvc": values.cvc
                        },
                    "encryption_key": "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a"
                },
            )

            if (resp.status == 200) {
                toast({
                    title: "Card Added Successfully",
                    status: "success",
                    isClosable: true,
                })

                onEncryptedCardChange(base64ToHex(resp.data.encrypted_card))
                onEncryptedKeyChange(resp.data.encryption_key)
                closeModal()
            }
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

    const {
        handleSubmit,
        register,
        formState: { errors, isSubmitting },
    } = useForm();

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
                >Add Card To Wallet</ModalHeader>
                <ModalCloseButton/>
                <form onSubmit={handleSubmit(addCard)}>
                <ModalBody>
                    <Grid templateColumns={"repeat(1, 1fr)"} gap={8} w={"full"}>
                        <GridItem
                            p={"2rem"}
                            borderRadius={"md"}
                        >
                            <Stack>
                                <FormControl id="card_number" isRequired isInvalid={errors.card_number}>
                                    <Input
                                        placeholder={"Card Number"}
                                        type="number"
                                        {...register("card_number", {
                                            required: "This is required",
                                            maxLength: {
                                                value: 16,
                                                message: "Invalid card number",
                                            },
                                            minLength: {
                                                value: 16,
                                                message: "Invalid card number",
                                            }
                                        })}
                                    />
                                    <FormErrorMessage>
                                        {errors.card_number && errors.card_number.message}
                                    </FormErrorMessage>
                                </FormControl>
                                <FormControl id="holder" isRequired isInvalid={errors.holder}>
                                    <Input
                                        placeholder={"Card Holder"}
                                        type="text"
                                        {...register("holder", {
                                            required: "This is required",
                                            minLength: {
                                                value: 4,
                                                message: "Minimum length should be 4",
                                            },
                                        })}
                                    />
                                    <FormErrorMessage>
                                        {errors.holder && errors.holder.message}
                                    </FormErrorMessage>
                                </FormControl>

                                <HStack>
                                    <FormControl id="exp_date" isRequired isInvalid={errors.exp_date}>
                                        <Input
                                            placeholder={"Exp. date"}
                                            type="text"
                                            {...register("exp_date", {
                                                required: "This is required",
                                                minLength: {
                                                    value: 5,
                                                    message: "Invalid expiration date",
                                                },
                                            })}
                                        />
                                        <FormErrorMessage>
                                            {errors.exp_date && errors.exp_date.message}
                                        </FormErrorMessage>
                                    </FormControl>
                                    <FormControl id="cvc" isRequired isInvalid={errors.cvc}>
                                        <Input
                                            placeholder={"CVC"}
                                            type="password"
                                            {...register("cvc", {
                                                required: "This is required",
                                                maxLength: {
                                                    value: 3,
                                                    message: "Invalid CVC",
                                                },
                                                minLength: {
                                                    value: 3,
                                                    message: "Invalid CVC",
                                                },
                                            })}
                                        />
                                        <FormErrorMessage>
                                            {errors.cvc && errors.cvc.message}
                                        </FormErrorMessage>
                                    </FormControl>
                                </HStack>
                            </Stack>

                        </GridItem>
                    </ Grid>
                </ModalBody>

                <ModalFooter
                    borderTopWidth='1px'
                    borderColor={"brand.200"}
                >
                    <Button
                        bg={"brand.200"}
                        color={"brand.100"}
                        _hover={
                            {bg: "brand.hover"}
                        }
                        type="submit"
                    >Add Card</Button>
                </ModalFooter>
                </form>
            </ModalContent>
        </Modal>
    )
};

export default CardModal;
