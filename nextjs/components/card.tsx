import { Flex, Icon, Stack, Text, useDisclosure } from "@chakra-ui/react";
import { PropsWithChildren } from "react";
import TransactionModal from "./transactionModal";
import { MdOutlineMoreVert } from "react-icons/md";
import { KeyedMutator } from "swr";

type Props = {
  token: string;
  mutator: KeyedMutator<any>;
  transaction: {
    id: string;
    status: string;
    payment_method_type: string;
    payment_method: object;
    amount: number;
    billing_address: {
      first_name: string;
      last_name: string;
      postal_code: string;
      city: string;
      address_line_1: string;
      email: string;
      phone: string;
    };
    created_at: string;
    updated_at: string;
  };
  theurl: string;
} & Record<string, any>;

const Card = ({ token, mutator, transaction, theurl }: PropsWithChildren<Props>) => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Flex
        maxH={"80px"}
        paddingRight={"1rem"}
        direction={"row"}
        alignItems={"center"}
        bg={"brand.300"}
        color={"brand.100"}
        borderLeft={"10px solid"}
        borderColor={statusColor(transaction.status)}
        borderRadius={"md"}
        cursor={"pointer"}
        _hover={{ bg: "brand.hover" }}
        onClick={onOpen}
      >
        <Stack
          direction={"column"}
          bg={"transparent"}
          padding={"1rem"}
          align={"flex-start"}
          w={"100%"}
        >
          <Text fontWeight={"bold"}>
            {typeFormat(transaction.payment_method_type)}
          </Text>
          <Text fontSize={"12px"}>{readableDate(transaction.created_at)}</Text>
        </Stack>

        <Icon as={MdOutlineMoreVert} w={6} h={6} color="brand.200" />
      </Flex>

      <TransactionModal
        key={transaction.id}
        token={token}
        mutator={mutator}
        isOpened={isOpen}
        closeModal={onClose}
        transaction={transaction}
        theurl={theurl}
      />
    </>
  );
};

export default Card;

function statusColor(status: string) {
  let color;

  switch (status) {
    case "pending_settlement":
      color = "brand.pending";
      break;
    case "settled":
      color = "brand.settled";
      break;
    case "rejected":
      color = "brand.rejected";
      break;
    default:
      color = "I have never heard of that fruit...";
  }

  return color;
}

function typeFormat(type: string) {
  let formattedType;

  switch (type) {
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

  return formattedType;
}

function readableDate(date: string) {
  const d = new Date(date);
  return d.toDateString();
}
