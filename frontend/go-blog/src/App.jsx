import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import chakra from './assets/chakra.png'
import './App.css'
import {Box, Heading, Stack,} from "@chakra-ui/react";

function App() {
  const [count, setCount] = useState(0)

  return (
    <Stack direction='column'>
        <Stack direction='row' spacing='24px' display='flex' alignItems='center' justifyContent='center'>
            <img src={reactLogo} alt="react logo" className="App-logo" />
            <img src={viteLogo} alt="vite logo" className="App-logo" />
            <img src={chakra} alt="chakra logo" className="App-logo"  width={40} style={{borderRadius:'50%' , overflow:'hidden'}} />
        </Stack>
        <Heading as='h1'size={'xl'}>Vite + React  + Chakra Ui</Heading>
    </Stack>
  )
}

export default App
