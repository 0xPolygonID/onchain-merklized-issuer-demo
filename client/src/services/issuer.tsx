import axios from 'axios';

const OnchainIssuerNodeHost = process.env.NEXT_PUBLIC_ISSUER_URL || 'http://localhost:8080';

interface ApiError extends Error {
    response?: {
        status: number;
    };
}

export async function getIssuersList(): Promise<string[]> {
    try {
        const response = await axios.get<string[]>(`${OnchainIssuerNodeHost}/api/v1/issuers`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

interface AuthQRCodeResponse {
    data: any;
    sessionId: string;
}

export async function produceAuthQRCode(issuer: string): Promise<AuthQRCodeResponse> {
    try {
        if (!issuer) {
            throw new Error('Issuer is not defined');
        }
        const url = new URL(`${OnchainIssuerNodeHost}/api/v1/requests/auth`);
        url.search = new URLSearchParams({ issuer: issuer }).toString();
        const response = await axios.get<any>(url.toString());
        return {
            data: response.data,
            sessionId: response.headers['x-id'],
        };
    } catch (error) {
        throw error;
    }
}

interface AuthSessionStatusResponse {
    id: string;
}

export async function checkAuthSessionStatus(sessionId: string): Promise<AuthSessionStatusResponse | null> {
    try {
        const url = new URL(`${OnchainIssuerNodeHost}/api/v1/status`);
        url.search = new URLSearchParams({ id: sessionId }).toString();
        const response = await axios.get<any>(url.toString());
        return {
            id: response.data.id,
        };
    } catch (error) {
        const apiError = error as ApiError;
        if (apiError.response && apiError.response.status === 404) {
            return null;
        }
        throw error;
    }
}

export interface CredentialRequest {
    credentialSchema: string;
    type: string;
    credentialSubject: {
        id: string;
        balance: number;
    };
    expiration: number;
}
export class CredentialBalanceRequest {
    private credentialSchema: string = 'ipfs://QmY3nuhAYYH13wDxNMzzqhSHpjeCz1pDYY7efzF2162xFE';
    private type: string = 'BalanceCredential';
    private id: string;
    private balance: number;
    private expiration: number;

    constructor(id: string, balance: BigInt, expiration?: number) {
        this.id = id;
        this.balance = Number(balance);
        this.expiration = expiration || Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 90;
    }

    construct(): CredentialRequest {
        return {
            credentialSchema: this.credentialSchema,
            type: this.type,
            credentialSubject: {
                id: this.id,
                balance: this.balance,
            },
            expiration: this.expiration,
        };
    }
}


interface NewCredentialResponse {
    id: string;
}

export async function requestIssueNewCredential(issuerDid: string, credentialRequest: CredentialRequest): Promise<NewCredentialResponse> {
    try {
        const response = await axios.post<any>(
            `${OnchainIssuerNodeHost}/api/v1/identities/${issuerDid}/claims`,
            credentialRequest // JSON payload
        );
        return {
            id: response.data.id,
        };
    } catch (error) {
        throw error;
    }
}


export async function getCredentialOffer(issuerDid: string, subjectDid: string, claimId: string): Promise<any> {
    try {
        const response = await axios.get<any>(
            `${OnchainIssuerNodeHost}/api/v1/identities/${issuerDid}/claims/offer?subject=${subjectDid}&claimId=${claimId}`
        );
        return response.data;
    } catch (error) {
        throw error;
    }
}