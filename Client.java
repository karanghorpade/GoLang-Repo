package com.company;

import java.io.*;
import java.net.Socket;

public class Client
{
    public static void main(String[] args) throws Exception
    {
        Socket s = new Socket("127.0.0.1", 3000);

        BufferedReader breader = new BufferedReader(new InputStreamReader(System.in));
        OutputStream os = s.getOutputStream();
        PrintWriter pw = new PrintWriter(os, true);

        InputStream is = s.getInputStream();
        BufferedReader r = new BufferedReader(new InputStreamReader(is));
        System.out.println("Type your message and press enter !!");
        System.out.println("Start Chatting - ");

        String receive, send;
        while(true)
        {
            send = breader.readLine();
            pw.println(send);
            pw.flush();
            if((receive = r.readLine()) != null)
            {
                System.out.println("Server: " +receive);
            }
        }
    }
}