import {
    Alert,
    AlertIcon, Box,
    Button,
    FormControl,
    FormErrorMessage,
    FormLabel, Grid, GridItem, HStack,
    Input,
    Stack,
    useToast,
} from "@chakra-ui/react";
import {PropsWithChildren} from "react";
import {useForm} from "react-hook-form";
import axios from "axios";

type Props = {
    token: string,
} & Record<string, any>;

const CreditCard = ( {token}: PropsWithChildren<Props>) => {

    const toast = useToast()

    const setAuthHeader = () => {
        return { headers: { Authorization: token } }
    }

    const onSubmit = async (values: any | undefined) => {
        try {
            const resp = await axios.post(
                'http://localhost:8080/api/transaction/creditcard',
                {
                    "amount": values.amount,
                    "payment_method": {
                        "credit_card": {
                            "card_number": values.card_number,
                            "holder_name": values.holder_name,
                            "exp_date": values.exp_date,
                            "cvc": values.cvc
                        }
                    },
                    "billing_address": {
                        "first_name": values.first_name,
                        "last_name": values.last_name,
                        "postal_code": values.postal_code,
                        "city": values.city,
                        "address_line_1": values.address_line_1,
                        "email": values.email,
                        "phone": values.phone
                    }
                },
                setAuthHeader()
            )
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
        <form onSubmit={handleSubmit(onSubmit)}>
            <Grid templateColumns={"repeat(2, 1fr)"} gap={8} w={"full"}>
                <GridItem
                    p={"2rem"}
                    bg={"brand.200"}
                    borderRadius={"md"}
                >
                    <Stack>
                        <FormControl id="amount" isInvalid={errors.amount}>
                            <FormLabel>Amount</FormLabel>
                            <Input
                                disabled
                                value={Math.trunc(Math.random()*2000) + 1}
                                type="number"
                                {...register("amount", )}
                            />
                            <FormErrorMessage>
                                {errors.amount && errors.amount.message}
                            </FormErrorMessage>
                        </FormControl>
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
                <GridItem
                    p={"2rem"}
                    bg={"brand.200"}
                    borderRadius={"md"}
                >
                    <Stack>
                        <HStack>
                            <FormControl id="first_name" isRequired isInvalid={errors.first_name}>
                                <Input
                                    placeholder={"First Name"}
                                    type="text"
                                    {...register("first_name", {
                                        required: "This is required",
                                        minLength: {
                                            value: 3,
                                            message: "Invalid First Name",
                                        },
                                    })}
                                />
                                <FormErrorMessage>
                                    {errors.first_name && errors.first_name.message}
                                </FormErrorMessage>
                            </FormControl>
                            <FormControl id="last_name" isRequired isInvalid={errors.last_name}>
                                <Input
                                    placeholder={"Last Name"}
                                    type="text"
                                    {...register("last_name", {
                                        required: "This is required",
                                        minLength: {
                                            value: 3,
                                            message: "Invalid Last Name",
                                        },
                                    })}
                                />
                                <FormErrorMessage>
                                    {errors.last_name && errors.last_name.message}
                                </FormErrorMessage>
                            </FormControl>

                        </HStack>
                        <HStack>
                            <FormControl id="postal_code" isRequired isInvalid={errors.postal_code}>
                                <Input
                                    placeholder={"Postal Code"}
                                    type="number"
                                    {...register("postal_code", {
                                        required: "This is required",
                                        minLength: {
                                            value: 2,
                                            message: "Invalid Postal Code",
                                        },
                                    })}
                                />
                                <FormErrorMessage>
                                    {errors.postal_code && errors.postal_code.message}
                                </FormErrorMessage>
                            </FormControl>
                            <FormControl id="city" isRequired isInvalid={errors.city}>
                                <Input
                                    placeholder={"City"}
                                    type="text"
                                    {...register("city", {
                                        required: "This is required",
                                        minLength: {
                                            value: 2,
                                            message: "Invalid City",
                                        },
                                    })}
                                />
                                <FormErrorMessage>
                                    {errors.city && errors.city.message}
                                </FormErrorMessage>
                            </FormControl>
                        </HStack>

                        <FormControl id="address_line_1" isRequired isInvalid={errors.address_line_1}>
                            <Input
                                placeholder={"Address Line"}
                                type="text"
                                {...register("address_line_1", {
                                    required: "This is required",
                                    minLength: {
                                        value: 2,
                                        message: "Invalid Address Line 1",
                                    },
                                })}
                            />
                            <FormErrorMessage>
                                {errors.address_line_1 && errors.address_line_1.message}
                            </FormErrorMessage>
                        </FormControl>

                        <HStack>
                            <FormControl id="phone" isRequired isInvalid={errors.phone}>
                                <Input
                                    placeholder={"Phone"}
                                    type="text"
                                    {...register("phone", {
                                        required: "This is required",
                                        minLength: {
                                            value: 2,
                                            message: "Invalid Phone Number",
                                        },
                                    })}
                                />
                                <FormErrorMessage>
                                    {errors.phone && errors.phone.message}
                                </FormErrorMessage>
                            </FormControl>
                            <FormControl id="email" isRequired isInvalid={errors.email}>
                                <Input
                                    placeholder={"Email"}
                                    type="text"
                                    {...register("email", {
                                        required: "This is required",
                                        minLength: {
                                            value: 2,
                                            message: "Invalid Email",
                                        },
                                    })}
                                />
                                <FormErrorMessage>
                                    {errors.email && errors.email.message}
                                </FormErrorMessage>
                            </FormControl>

                        </HStack>
                    </Stack>

                </GridItem>
                <GridItem
                    colSpan={2}
                    justifySelf={"end"}
                >
                    <Button
                        isLoading={isSubmitting}
                        loadingText="Process Transaction"
                        type="submit"
                        bg={"brand.200"}
                        color={"brand.100"}
                        _hover={
                            {bg: "brand.hover"}
                        }
                    >
                        Process Transaction
                    </Button>
                </GridItem>
            </Grid>

        </form>
    )
};

export default CreditCard;
