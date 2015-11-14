/**
 * Prints 'Hello, world!' in a variety of ways and demonstrates basic
 * java knowledge
 */


import skilstak.c;

public class hello {
/**
 * Functions to print a message, typically 'Hello, world!' in
 * a variety of ways.
 */
    /**
     * Prints the message repetedly in a number of colors in
     * a 'nyan'-esque style.
     */
    public static void color(String message) {
        System.out.println(c.clear);
        while (true) {
            System.out.print(c.rc() + message + " ");
        }
    }
    /**
     * Prints a message with no fanfare.
     */
    public static void plain(String message) {
        System.out.println(c.clear + message);
    }

    /**
     * Prints a message, Las Vegas blinking sign style.
     */

    public static void multi(String message) throws InterruptedException {
        while (true) {
            System.out.print(c.clear + c.multi(message));
            Thread.sleep(500);
        }
    }

    /** Determines which arguments are options and which are a message */

    public static void main(String[] args) throws InterruptedException {
        String who = "world";
        String option = "";
        String message = "";

        if (args.length > 1) {
            option = args[0];
            who = args[1];
        } else if (args.length == 1) {
            if (args[0].startsWith("-")) {
                option = args[0];
            } else {
                who = args[0];
            }
        }
        
        message = "Hello " + who + "!";

        if (option.equals("-m")) {
            multi(message);
        } else if (option.equals("-c")) {
            color(message);
        } else {
            plain(message);
        }
    }
}
