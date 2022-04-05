import {
    Alert,
    AlertIcon, Box,
    Button,
    FormControl,
    FormErrorMessage,
    FormLabel, Grid, GridItem, HStack,
    Input, Select,
    Stack,
    useToast,
} from "@chakra-ui/react";
import {PropsWithChildren} from "react";
import {useForm} from "react-hook-form";
import axios from "axios";

type Props = {
    token: string,
} & Record<string, any>;

const ACH = ( {token}: PropsWithChildren<Props>) => {

    const toast = useToast()

    const setAuthHeader = () => {
        return { headers: { Authorization: token } }
    }

    const onSubmit = async (values: any | undefined) => {
        try {
            const resp = await axios.post(
                'http://localhost:8080/api/transaction/ach',
                {
                    "amount": values.amount,
                    "payment_method": {
                        "ach": {
                            "account_number": values.account_number,
                            "routing_number": values.routing_number,
                            "account_type": values.account_type,
                            "sec_code": values.sec_code
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
                        <FormControl id="account_number" isRequired isInvalid={errors.account_number}>
                            <Input
                                placeholder={"Account Number"}
                                type="number"
                                {...register("account_number", {
                                    required: "This is required",
                                })}
                            />
                            <FormErrorMessage>
                                {errors.account_number && errors.account_number.message}
                            </FormErrorMessage>
                        </FormControl>
                        <FormControl id="routing_number" isRequired isInvalid={errors.routing_number}>
                            <Input
                                placeholder={"Routing Number"}
                                type="text"
                                {...register("routing_number", {
                                    required: "This is required",
                                })}
                            />
                            <FormErrorMessage>
                                {errors.routing_number && errors.routing_number.message}
                            </FormErrorMessage>
                        </FormControl>

                        <HStack>
                            <FormControl id="account_type" isRequired isInvalid={errors.account_type}>
                                <Select placeholder='Account Type'
                                        {...register("account_type", {
                                            required: "This is required",
                                        })}
                                >
                                    <option value='checking'>Checking</option>
                                    <option value='savings'>Savings</option>
                                </Select>
                                <FormErrorMessage>
                                    {errors.account_type && errors.account_type.message}
                                </FormErrorMessage>
                            </FormControl>
                            <FormControl id="sec_code" isRequired isInvalid={errors.sec_code}>
                                <Input
                                    placeholder={"SEC Code"}
                                    type="password"
                                    {...register("sec_code", {
                                        required: "This is required",
                                    })}
                                />
                                <FormErrorMessage>
                                    {errors.sec_code && errors.sec_code.message}
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

export default ACH;
