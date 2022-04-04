import {ReactNode, useState} from 'react';
import {
    Box,
    Flex,
    Button,
} from '@chakra-ui/react';

const TokenCard = ({ children }: { children: ReactNode }) => (
    <Box
        px={'20px'}
        py={'10px'}
        rounded={'md'}
        bg={'gray.600'}
        color={'gray.50'}
        maxW={'300px'}
        overflow={'hidden'}
        whiteSpace={'nowrap'}
        textOverflow={'ellipsis'}
    >
        {children}
    </Box>
);

export default function Nav() {

    const [token, setToken] = useState()

    const getToken = async () => {
        const response = await fetch( "http://localhost:8080/login")
        const data = await response.json()
        setToken(data.token)
    }

    return (
        <>
            <Box
                bg={'gray.700'}
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
                        <TokenCard>{token}</TokenCard>
                        :
                        <Button onClick={getToken}>
                            Get Token
                        </Button>
                    }
                </Flex>
            </Box>
        </>
    );
}
