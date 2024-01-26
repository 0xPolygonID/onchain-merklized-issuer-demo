'use client'

import React, { useState, useContext, useEffect } from 'react';
import { useRouter } from 'next/router';
import { Grid, Box, Typography, Button, Backdrop, CircularProgress} from '@mui/material';
import JSONPretty from 'react-json-pretty';
import {selectMetamaskWallet, getMetamaskWalletBalance, FormatBalance} from '@/services/metamask';
import { requestIssueNewCredential } from '@/services/issuer'
import { CredentialBalanceRequest, CredentialRequest } from '@/services/issuer';
import { ErrorPopup } from '@/app/components';
import SelectedIssuerContext from '@/contexts/SelectedIssuerContext';

const App = () => {
  const router = useRouter();
  const routerQuery = router.query;
  

  const { selectedIssuerContext } = useContext(SelectedIssuerContext);
  useEffect(() => {
    if (!selectedIssuerContext) {
      router.push('/');
      return;
    }
  }, [selectedIssuerContext, router]);

  const [error, setError] = useState<string | null>(null);

  const [metamaskWalletAddress, setMetamaskwalletAddress] = useState('');
  const getMetamaskWallet = async () => {
    try {
      const wallet = await selectMetamaskWallet();
      setMetamaskwalletAddress(wallet.address);
    } catch (error) {
        setError(`Failed to get wallet address: ${error}`);
    }
  };

  const [credentialRequest, setCredentialRequest] = useState<CredentialRequest>();
  const retriveCredentialRequest = async () => {
    let gweiBalance: BigInt;
    try {
      gweiBalance = await getMetamaskWalletBalance(metamaskWalletAddress, FormatBalance.Gwei) 
      const cr = new CredentialBalanceRequest(routerQuery.userID as string, gweiBalance).construct();
      setCredentialRequest(cr); 
    } catch (error) {
      setError(`Failed to get wallet balance: ${error}`);
    }
  };

  const [isLoaded, setIsLoaded] = useState(false);
  const newCredentialRequest = async () => {
      setIsLoaded(true);
      try {
        if (!credentialRequest) {
          throw new Error('Credential request is not set');
        }

        const credentialInfo = await requestIssueNewCredential(selectedIssuerContext, credentialRequest)
        router.push(`/offer?claimId=${credentialInfo.id}&issuer=${selectedIssuerContext}&subject=${routerQuery.userID as string}`);
      } catch (error) {
        setError(`Error making the request: ${error}`);
      } finally {
        setIsLoaded(false);
      }
  }

  return (
    <Grid container 
      direction="column" 
      justifyContent="center" 
      alignItems="center"
      height="100%"
    >
    {error && <ErrorPopup error={error} />}
    
    {
      !metamaskWalletAddress && (
        <Box textAlign="center">
          <Typography variant="h6">
            Balance claim for user {routerQuery.userID}
          </Typography>
          <Button onClick={getMetamaskWallet} variant="contained" size="large">
            Connect MetaMask
          </Button>
        </Box>
      )
    }  

    {metamaskWalletAddress && !credentialRequest &&  (
      <Grid container direction="column" alignItems="center" textAlign="center">
        <Typography variant="h6">
          Account: {metamaskWalletAddress}
        </Typography>
        <Button onClick={retriveCredentialRequest} variant="contained" size="large">
          Get Balance GWEI
        </Button>
      </Grid>
    )}
  
    {credentialRequest && (
      <Grid container direction="column" alignItems="center" textAlign="center">
        <Typography variant="h6" textAlign="left">
            Credential request content:
        </Typography>
        <Grid textAlign="left">
          <Box alignItems="left">
            <JSONPretty 
            id="json-pretty" 
            style={{
              fontSize: "1.3em",
            }} 
            data={JSON.stringify(credentialRequest)}
            theme={jsonStyle}
          ></JSONPretty>
          </Box>
        </Grid>
        <Button onClick={newCredentialRequest} variant="contained" size="large">
          Get claim
        </Button>
      </Grid>
    )}
  
    <Backdrop
      sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1 }}
      open={isLoaded}
    >
      <CircularProgress color="inherit" />
    </Backdrop>
  </Grid>  
  );
};

const jsonStyle = {
  main: 'line-height:1.3;color:#66d9ef;background:#272822;overflow:auto;',
  error: 'line-height:1.3;color:#66d9ef;background:#272822;overflow:auto;',
  key: 'color:#f92672;',
  string: 'color:#fd971f;',
  value: 'color:#a6e22e;',
  boolean: 'color:#ac81fe;',
}

export default App;
