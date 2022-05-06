import type { NextPage } from "next";
import {
    Container,
    Flex,
    Text,
} from "@chakra-ui/react";

const Small: NextPage = (Props) => {

    return (
        <>
            <Container>
                <Flex
                    h={"600px"}
                    w={"full"}
                    justifyContent={"center"}
                    alignItems={"center"}
                >
                    <Text
                        color={"brand.100"}
                        fontWeight={"600"}
                    >
                        Please use bigger screen size.
                    </Text>
                </Flex>
            </Container>
        </>
    );
};

export default Small;
