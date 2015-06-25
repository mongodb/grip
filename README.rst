=======================================================
``grip`` -- A Go Library for Logging and Error Handling
=======================================================

*Under Construction*

``grip`` isn't any thing special, but it does a few pretty great
things:

#. Provide a common logging interface to log to both standard
   logging methods and/or ``systemd`` journal components.

#. Provides some simple methods for handling errors, particularly when
   you want to accumulate and then return errors.

*You just get a grip, folks.*

Use
---

Download:

::

   go get -u github.com/tychoish/grip

Import:

::

   import "github.com/tychoish/grip"

Components
----------

Logging
~~~~~~~

Provides a logging system that logs to ``systemd``'s ``journald``,
by default, falling back to standard output logging using the standard
library logging system. This facility is provided using the
``Journaler`` type and associated methods.

``systemd`` integration comes from
`go-systemd <https://github.com/coreos/go-systemd>`_, which is great.

By default ``grip.std`` defines a standard global  instances
that you can use with a set of ``grip.<Level>`` functions, or you can
create your own ``Journaler`` instance.

Defined helpers exist for the following levels/actions:

- ``Debug``
- ``Info``
- ``Notice``
- ``Warning``
- ``Error``
- ``ErrorPanic``
- ``ErrorFatal``
- ``Critical``
- ``CriticalPanic``
- ``CriticalFatal``
- ``Alert``
- ``Alert Panic``
- ``AlertFatal``
- ``Emergency``
- ``EmergencyPanic``
- ``EmergencyFatal``

Helpers ending with ``Panic`` call ``panic()`` after logging the message
message, and helpers ending with ``Fatal`` call ``OS.Exit(1)`` after
logging the message. Use responsibly.

``Journaler`` instances have a notion of "default" log levels and
thresholds, which provide the basis for verbosity control and sane
default behavior. The default level defines the priority/level of any
message with either an invalid priority specified *or* using the
``SendDefault`` helper. The threshold level, defines the minimum
priority or level that ``grip`` sends to the logging system. Consider
the following behaviors:

- It's possible to suppress message higher than the current default,
  such that default messages are never printed. This is not
  recommended.

- The helpers with ``Panic`` (message passed to panic) and ``Fatal``
  (messages ignored) helpers will still panic or exit even if their
  messages are suppressed.

- It's not possible to suppress the highest log level, ``Emergency``
  messages will always log.

``Journaler`` objects have the following, additional methods (also
available as functions in the ``grip`` package to manage the global
standard logger instance.):

- ``Send(<level int>, <message string>)`` to manually send a
  message. Levels are values between ``0`` and ``7``, where lower
  numbers are *more* severe.

- ``SendDefault(<message string>)`` to log using the default level.

- ``SetName(<string>)`` to reset the name of the logger and fallback
  logger instance. ``grip`` attempts to set this to the name of your
  program, but will fallback to ``go-grip-default-logger`` (typically
  when program is invoked with ``go run``.)

- ``SetFallback(<*log.Logger>)`` to pass your own logging instance to
  use when ``systemd`` is unavailable.

- ``SetDefault(<level int>)`` change the default log level. Levels are
  values between ``0`` and ``7``, where lower numbers are *more*
  severe. ``grip`` does *not* forbid configurations where default
  levels are *below* the configured threshold.

- ``SetThreshold(<level int>)`` Change the lowest log level that the
  ``grip`` will transmit to the logging mechanism (either ``systemd``
  ``journald`` or Go's standard logging.) Log messages with lower
  levels are not captured and ignored.

The ``Journaler.InvertFallback`` flag (bool) switches a ``Journaler``
instance to prefer the standard logging mechanism rather than
``systemd``.

By default:

- the log level uses the "Notice" level (``5``)

- the minimum threshold for logging is the "Info" level (``6``)
  (suppressing only debug.)

- fallback logging writes to standard output.

- ``InvertFanllback`` is ``false``, ``systemd`` logging is prefered.

Collector for "Continue on Error" Semantics
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

If you want to do something other than just swallow errors, but don't
need to hard abort, the ``MultiCatcher`` object makes this pattern
swell, a la:

::

   func doStuff(dirname string) (error) {
           files, err := ioutil.ReadDir(dirname)
           if err != nil {
                   // should abort here because we shouldn't continue.
                   return err
           }

           catcher := grip.NewCatcher()
           for _, f := range files {
               err = doStuffToFile(f.Name())
               catcher.Add(err)
           }

           return catcher.Resolve()
   }


Simple Error Catching
~~~~~~~~~~~~~~~~~~~~~

Use ``grip.Catch(<err>)`` to check and print error messages.

There are also helper functions on ``Journaler`` objects that check
and log error messages using either the default (global) ``Journaler``
instance, or a specific ``Journaler`` instance, at all levels.

- ``CatchDebug``
- ``CatchInfo``
- ``CatchNotice``
- ``CatchWarning``
- ``CatchError``
- ``CatchErrorPanic``
- ``CatchErrorFatal``
- ``CatchCritical``
- ``CatchCriticalPanic``
- ``CatchCriticalFatal``
- ``CatchAlert``
- ``CatchAlertPanic``
- ``CatchAlertFatal``
- ``CatchEmergency``
- ``CatchEmergencyPanic``
- ``CatchEmergencyFatal``
