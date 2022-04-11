import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Grid,
  GridItem,
  HStack,
  Input,
  Stack,
  useDisclosure,
  useToast,
} from "@chakra-ui/react";
import { PropsWithChildren, useState } from "react";
import { useForm } from "react-hook-form";
import axios from "axios";
import { KeyedMutator } from "swr";
import CardModal from "./cardModal";
import { MdDoneAll } from "react-icons/md";

type Props = {
  token: string;
  mutator: KeyedMutator<any>;
} & Record<string, any>;

const TokenPayment = ({ token, mutator, type }: PropsWithChildren<Props>) => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const toast = useToast();
  const setAuthHeader = () => {
    return { headers: { Authorization: token } };
  };
  const [encrypted, setEncrypted] = useState();
  const [identifier, setIdentifier] = useState();

  const onSubmitApplePay = async (values: any | undefined) => {
    try {
      const resp = await axios.post(
        "http://localhost:8080/api/transaction/applepay",
        {
          amount: Math.trunc(Math.random() * 2000) + 1,
          payment_method: {
            apple_pay: {
              payment_token: {
                identifier: identifier,
                payment_data: encrypted,
              },
            },
          },
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
      } else {
        toast({
          title: resp.data.message,
          status: "error",
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

  const onSubmitGooglePay = async (values: any | undefined) => {
    try {
      const resp = await axios.post(
        "http://localhost:8080/api/transaction/googlepay",
        {
          amount: 1099,
          payment_method: {
            google_pay: {
              encrypted_payment: {
                payment_id: identifier,
                payment_data: encrypted,
              },
            },
          },
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
      } else {
        toast({
          title: resp.data.message,
          status: "error",
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

  return (
    <>
      <CardModal
        isOpened={isOpen}
        closeModal={onClose}
        onEncryptedCardChange={setEncrypted}
        onEncryptedKeyChange={setIdentifier}
      />
      <form
        onSubmit={
          type == "applepay"
            ? handleSubmit(onSubmitApplePay)
            : handleSubmit(onSubmitGooglePay)
        }
      >
        <Grid templateColumns={"repeat(2, 1fr)"} gap={8} w={"full"}>
          <GridItem p={"2rem"} bg={"brand.200"} borderRadius={"md"}>
            <Flex
              w={"full"}
              justifyContent={"flex-start"}
              alignItems={"center"}
            >
              <Button
                bg={"brand.300"}
                color={"brand.100"}
                _hover={{ bg: "brand.hover" }}
                onClick={onOpen}
              >
                Add Card to Wallet
              </Button>

              {encrypted && identifier ? (
                <Button
                  rightIcon={<MdDoneAll />}
                  bg={"brand.200"}
                  color={"brand.settled"}
                  _hover={{ pointerEvents: "none" }}
                  cursor={"default"}
                >
                  Card Added
                </Button>
              ) : (
                <></>
              )}
            </Flex>
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
            {encrypted && identifier ? (
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
            ) : (
              <Button disabled={true} bg={"brand.200"} color={"brand.100"}>
                Add Card First
              </Button>
            )}
          </GridItem>
        </Grid>
      </form>
    </>
  );
};

export default TokenPayment;
