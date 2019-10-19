package com.company;
import java.io.*;
import java.net.*;
public class Server
{
    public static void main(String[] args) throws Exception
    {
        ServerSocket ss = new ServerSocket(3000);
        System.out.println("Server ready to chat with client - ");
        Socket sock = ss.accept( );
        // reading from keyboard (keyRead object)
        BufferedReader read = new BufferedReader(new InputStreamReader(System.in));
        // sending to client (pwrite object)
        OutputStream os = sock.getOutputStream();
        PrintWriter pw = new PrintWriter(os, true);

        // receiving from server ( receiveRead  object)
        InputStream is = sock.getInputStream();
        BufferedReader rr = new BufferedReader(new InputStreamReader(is));

        String receive, send;
        while(true)
        {
            if((receive = rr.readLine()) != null)
            {
                System.out.println("Client: "+receive);
            }
            send = read.readLine();
            pw.println(send);
            pw.flush();
        }
    }
}
