import Web3 from 'web3';
import {fromWei} from 'web3-utils';

declare global {
    interface Window {
        ethereum?: any;
    }
}

interface MetamaskError {
    message: string;
    code: number;
}

interface MetamaskWalletResponse {
    address: string;
}

export async function selectMetamaskWallet(): Promise<MetamaskWalletResponse> {
    try {
        const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
        if (!accounts || !accounts.length) {
            throw new Error('No accounts found');
        }
        return {
            address: accounts[0],
        };
    } catch (error) {
        const metamaskError = error as MetamaskError;
        // user rejected request
        if (metamaskError.code === 4001) {
            throw new Error('User rejected request');
        }
        throw error;
    }
}


export enum FormatBalance {
    Gwei = 'gwei',
    Wei = 'wei',
}

export async function getMetamaskWalletBalance(address: string, balanceFormat: FormatBalance = FormatBalance.Wei): Promise<BigInt> {
    const weiBalance = await (new Web3(window.ethereum)).eth.getBalance(address);
    if (balanceFormat === FormatBalance.Wei) {
        return weiBalance;
    }
    return BigInt(Math.floor(parseFloat(fromWei(weiBalance, 'gwei'))));
}