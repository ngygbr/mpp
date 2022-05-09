import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Grid,
  GridItem,
  HStack,
  Input, InputGroup, InputLeftAddon,
  Stack, Text,
  useToast,
} from "@chakra-ui/react";
import {PropsWithChildren, useState} from "react";
import { useForm } from "react-hook-form";
import axios from "axios";
import { KeyedMutator } from "swr";

type Props = {
  token: string;
  mutator: KeyedMutator<any>;
  theurl: string
} & Record<string, any>;

const CreditCard = ({ token, mutator, theurl }: PropsWithChildren<Props>) => {
  const toast = useToast();

  const setAuthHeader = () => {
    return { headers: { Authorization: token } };
  };

  const onSubmit = async (values: any | undefined) => {
    try {
      const resp = await axios.post(
        `${theurl}/api/transaction/creditcard`,
        {
          payment_method: {
            credit_card: {
              card_number: values.card_number,
              holder_name: values.holder_name,
              exp_date: values.exp_date,
              cvc: values.cvc,
            },
          },
          amount: parseInt(values.amount),
          billing_address: {
            first_name: values.first_name,
            last_name: values.last_name,
            postal_code: values.postal_code,
            city: values.city,
            address_line_1: values.address_line_1,
            email: values.email,
            phone: values.phone,
          },
        },
        setAuthHeader()
      );
      if (resp.status == 200) {
        toast({
          title: resp.data.message,
          status: "success",
          isClosable: true,
        });
      }
      await mutator(token);
    } catch (error: any) {
      if (error.response) {
        toast({
          title: error.response.data.message,
          status: "error",
          isClosable: true,
        });
      } else {
        toast({
          title: error.message,
          status: "error",
          isClosable: true,
        });
      }
    }
  };

  const {
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useForm();

  // @ts-ignore
  // @ts-ignore
  // @ts-ignore
  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Grid templateColumns={"repeat(2, 1fr)"} gap={8} w={"full"}>
        <GridItem p={"2rem"} bg={"brand.200"} borderRadius={"md"}>
          <Stack>
            <FormControl id="amount" isInvalid={errors.amount}>
              <InputGroup>
              <InputLeftAddon
                  bg={"brand.300"}
                  color={"brand.100"}
                  border={"none"}
              ><Text>Amount</Text></InputLeftAddon>
              <Input
                _hover={{ pointerEvents: "none" }}
                cursor={"default"}
                value={Math.trunc(Math.random() * 2000) + 1}
                type="number"
                {...register("amount")}
              />
              </InputGroup>
              <FormErrorMessage>
                {errors.amount && errors.amount.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl
              id="card_number"
              isRequired
              isInvalid={errors.card_number}
            >
              <InputGroup>
                <InputLeftAddon
                    bg={"brand.300"}
                    color={"brand.100"}
                    border={"none"}
                ><Text>Card Number</Text></InputLeftAddon>
              <Input
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
                  },
                })}
              />
              </InputGroup>
              <FormErrorMessage>
                {errors.card_number && errors.card_number.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl id="holder_name" isRequired isInvalid={errors.holder_name}>
              <InputGroup>
                <InputLeftAddon
                    bg={"brand.300"}
                    color={"brand.100"}
                    border={"none"}
                ><Text>Card Holder</Text></InputLeftAddon>
              <Input
                type="text"
                {...register("holder_name", {
                  required: "This is required",
                  minLength: {
                    value: 4,
                    message: "Minimum length should be 4",
                  },
                })}
              />
              </InputGroup>
              <FormErrorMessage>
                {errors.holder_name && errors.holder_name.message}
              </FormErrorMessage>
            </FormControl>

            <HStack>
              <FormControl id="exp_date" isRequired isInvalid={errors.exp_date}>
                <InputGroup>
                  <InputLeftAddon
                      bg={"brand.300"}
                      color={"brand.100"}
                      border={"none"}
                  ><Text>Exp. Date</Text></InputLeftAddon>
                <Input
                  placeholder={"MM/YY"}
                  type="text"
                  {...register("exp_date", {
                    required: "This is required",
                    minLength: {
                      value: 5,
                      message: "Invalid expiration date",
                    },
                  })}
                />
                </InputGroup>
                <FormErrorMessage>
                  {errors.exp_date && errors.exp_date.message}
                </FormErrorMessage>
              </FormControl>
              <FormControl id="cvc" isRequired isInvalid={errors.cvc}>
                <InputGroup>
                  <InputLeftAddon
                      bg={"brand.300"}
                      color={"brand.100"}
                      border={"none"}
                  ><Text>CVC</Text></InputLeftAddon>
                <Input
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
                </InputGroup>
                <FormErrorMessage>
                  {errors.cvc && errors.cvc.message}
                </FormErrorMessage>
              </FormControl>
            </HStack>
          </Stack>
        </GridItem>
        <GridItem p={"2rem"} bg={"brand.200"} borderRadius={"md"}>
          <Stack>
            <HStack>
              <FormControl
                id="first_name"
                isRequired
                isInvalid={errors.first_name}
              >
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
              <FormControl
                id="last_name"
                isRequired
                isInvalid={errors.last_name}
              >
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
              <FormControl
                id="postal_code"
                isRequired
                isInvalid={errors.postal_code}
              >
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

            <FormControl
              id="address_line_1"
              isRequired
              isInvalid={errors.address_line_1}
            >
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
        <GridItem colSpan={2} justifySelf={"end"}>
          <Button
            isLoading={isSubmitting}
            loadingText="Process Transaction"
            type="submit"
            bg={"brand.200"}
            color={"brand.100"}
            _hover={{ bg: "brand.hover" }}
          >
            Process Transaction
          </Button>
        </GridItem>
      </Grid>
    </form>
  );
};

export default CreditCard;
