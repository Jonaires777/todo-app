import { BASE_API_URL } from "@/utils/constants";
import axios, { AxiosRequestConfig } from "axios";

class ApiProvider {
  private instance;

  constructor() {
    this.instance = axios.create({
      baseURL: BASE_API_URL,
    });

    this.instance.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem("token");
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }

        return config;
      },
      (error) => {
        return Promise.reject(error);
      }
    );
  }

  async get<T>(path: string, options?: AxiosRequestConfig): Promise<T> {
    try {
      const response = await this.instance.get<T>(path, options);

      if (response.status < 200 || response.status > 300) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async post<T, D>(
    path: string,
    data?: D,
    options?: AxiosRequestConfig
  ): Promise<T> {
    try {
      const response = await this.instance.post<T>(path, data, options);

      if (response.status < 200 || response.status > 300) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async put<T, D>(
    path: string,
    data?: D,
    options?: AxiosRequestConfig
  ): Promise<T> {
    try {
      const response = await this.instance.put<T>(path, data, options);

      if (response.status < 200 || response.status > 300) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async delete<T>(path: string, options?: AxiosRequestConfig): Promise<T> {
    try {
      const response = await this.instance.delete<T>(path, options);

      if (response.status < 200 || response.status > 300) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      return response.data;
    } catch (error) {
      throw error;
    }
  }
}

const apiProvider = new ApiProvider();

export default apiProvider;
