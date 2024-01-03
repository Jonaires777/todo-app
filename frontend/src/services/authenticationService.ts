import apiProvider from "@/providers/todoProvider";

type AuthenticationResponse = {
  token: string;
};

type AuthenticationRequest = {
  email: string;
  password: string;
};

type SignUpResponse = {
  token: string;
};

type SignUpRequest = {
  name: string;
  email: string;
  password: string;
};

const login = async (
  email: string,
  password: string
): Promise<AuthenticationResponse> => {
  try {
    const response = await apiProvider.post<
      AuthenticationResponse,
      AuthenticationRequest
    >("/api/auth/login", { email, password });
    return response;
  } catch (error) {
    throw error;
  }
};

const signUp = async (
  email: string,
  name: string,
  password: string
): Promise<SignUpResponse> => {
  try {
    const signUpResponse = await apiProvider.post<
      SignUpResponse,
      SignUpRequest
    >("/api/auth/signup", {
      name,
      email,
      password,
    });
    return signUpResponse;
  } catch (error) {
    throw error;
  }
};

const checkIfAuthenticated = async (): Promise<boolean> => {
  try {
    await apiProvider.get("/api/auth/authorized-only");
    return true;
  } catch (error) {
    return false;
  }
};

const authenticationService = {
  login,
  signUp,
  checkIfAuthenticated,
};

export default authenticationService;
