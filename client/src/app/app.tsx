import { MantineProvider } from '@mantine/core'
import type { AppProps } from 'next/app'

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <MantineProvider>
      {/* <Layout> */}
      <Component {...pageProps} />
      {/* </Layout> */}
    </MantineProvider>
  )
}

export default App
