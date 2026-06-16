# SpellingGopher — Architecture Diagrams
---

## 1. Package Dependencies & Layers

How the packages wire together, grouped by architectural layer. Arrows point in
the direction of the dependency (caller → callee).

```mermaid
%%{init: {'theme':'base', 'themeVariables': {
  'fontFamily':'ui-monospace, SFMono-Regular, Menlo, monospace',
  'fontSize':'15px',
  'primaryColor':'#1f2430',
  'primaryTextColor':'#e6e6e6',
  'primaryBorderColor':'#5ccfe6',
  'lineColor':'#8a9199',
  'clusterBkg':'#161821',
  'clusterBorder':'#2d3343'
}}}%%
flowchart TB
    subgraph ENTRY["🚪 Entrypoint"]
        MAIN["cmd/spelling-gopher<br/><b>main</b>"]
    end

    subgraph UI["🖥️ Presentation · TUI"]
        TUI["internal/tui<br/><b>Bubble Tea Model</b>"]
    end

    subgraph DOMAIN["🎯 Domain Core"]
        QUOTE["internal/quote<br/><b>Quote · Service · Repository</b>"]
        GAME["internal/game<br/><b>Session · Stats · Glyph · Clock</b>"]
    end

    subgraph INFRA["🔌 Infrastructure · Adapters"]
        CSV["internal/infra/csvquotes<br/><b>Repository</b>"]
        ZEN["internal/infra/zenquotes<br/><b>Repository</b>"]
    end

    MAIN --> TUI
    MAIN --> QUOTE
    MAIN --> GAME
    MAIN --> CSV

    TUI --> GAME
    TUI --> QUOTE

    CSV -. implements .-> QUOTE
    ZEN -. implements .-> QUOTE

    classDef entry  fill:#2b2233,stroke:#d4bfff,stroke-width:2px,color:#f2e9ff;
    classDef ui     fill:#1e2a33,stroke:#5ccfe6,stroke-width:2px,color:#d6f5ff;
    classDef domain fill:#26221a,stroke:#ffcc66,stroke-width:2px,color:#fff3d6;
    classDef infra  fill:#1f2a22,stroke:#bae67e,stroke-width:2px,color:#e8ffd6;

    class MAIN entry;
    class TUI ui;
    class QUOTE,GAME domain;
    class CSV,ZEN infra;
```

> The **dotted arrows** are the key to the design: infrastructure adapters
> *implement* the `quote.Repository` interface, so the domain never depends on
> a concrete data source. Swapping CSV for the ZenQuotes HTTP API is a one-line
> change in `main`.

---

## 2. Types & Relationships (Class Diagram)

The concrete structs, the interfaces they satisfy, and how they compose.

```mermaid
%%{init: {'theme':'base', 'themeVariables': {
  'fontFamily':'ui-monospace, SFMono-Regular, Menlo, monospace',
  'fontSize':'14px',
  'primaryColor':'#1f2430',
  'primaryTextColor':'#e6e6e6',
  'primaryBorderColor':'#5ccfe6',
  'lineColor':'#8a9199'
}}}%%
classDiagram
    direction LR

    class Repository {
        <<interface>>
        +Random(ctx) Quote, error
    }

    class Quote {
        +string Text
        +string Author
        +RuneCount() int
    }

    class Service {
        -Repository repo
        +NewService(r) Service
        +Random(ctx) Quote, error
    }

    class Clock {
        <<interface>>
        +Now() Time
    }

    class RealClock {
        +Now() Time
    }

    class Session {
        -[]rune target
        -[]rune typed
        +string Author
        -int keystrokes
        -int errors
        -Clock clock
        +Type(r rune)
        +Backspace()
        +Glyphs() []Glyph
        +Stats() Stats
        +Finished() bool
    }

    class Glyph {
        +rune Expected
        +rune Current
        +CharState State
        +IsSpace() bool
    }

    class Stats {
        +Duration Elapsed
        +float64 WPM
        +float64 Accuracy
    }

    class CsvRepository {
        -[]Quote quotes
        +Random(ctx) Quote, error
    }

    class ZenRepository {
        -http.Client client
        -string baseURL
        +Random(ctx) Quote, error
    }

    Service o-- Repository : depends on
    Service ..> Quote : returns
    Repository ..> Quote : returns
    CsvRepository ..|> Repository : implements
    ZenRepository ..|> Repository : implements
    Session o-- Clock : uses
    RealClock ..|> Clock : implements
    Session ..> Glyph : produces
    Session ..> Stats : produces
    Glyph --> CharState : has
```

---

## 3. Runtime Flow — A Full Typing Round

The Bubble Tea event loop, from launch to results and back. Solid lines are
commands/calls; dashed lines are messages flowing back into `Update`.

```mermaid
%%{init: {'theme':'base', 'themeVariables': {
  'fontFamily':'ui-monospace, SFMono-Regular, Menlo, monospace',
  'fontSize':'14px',
  'primaryColor':'#1f2430',
  'primaryTextColor':'#e6e6e6',
  'primaryBorderColor':'#5ccfe6',
  'actorBkg':'#1e2a33',
  'actorBorder':'#5ccfe6',
  'actorTextColor':'#d6f5ff',
  'signalColor':'#8a9199',
  'signalTextColor':'#e6e6e6',
  'noteBkgColor':'#26221a',
  'noteBorderColor':'#ffcc66',
  'noteTextColor':'#fff3d6'
}}}%%
sequenceDiagram
    autonumber
    actor U as 👤 User
    participant M as Model (tui)
    participant T as typingModel
    participant Svc as quote.Service
    participant Repo as Repository
    participant S as game.Session
    participant R as resultsModel

    M->>T: Init()
    T->>Svc: fetchQuote → Random(ctx)
    Svc->>Repo: Random(ctx)
    Repo-->>T: quoteMsg{quote}
    Note over T: NewSession(text, clock, author)

    loop Each keypress
        U->>M: KeyPressMsg
        M->>T: Update(msg)
        T->>S: Type(r) / Backspace()
        S-->>T: Glyphs() + Stats()
        Note over T,U: render colored glyphs + live WPM
    end

    S-->>T: Finished() == true
    T-->>M: sessionFinishedMsg{session}
    M->>R: active = screenResults
    R-->>U: show WPM · accuracy · time

    U->>M: enter (restart)
    M-->>M: restartGameMsg → reset & Init()
    U->>M: esc / ctrl+c
    M-->>U: tea.Quit
```

---

## 4. Screen State Machine

The TUI is a two-state machine switched by messages inside `Model.Update`.

```mermaid
%%{init: {'theme':'base', 'themeVariables': {
  'fontFamily':'ui-monospace, SFMono-Regular, Menlo, monospace',
  'fontSize':'15px',
  'primaryColor':'#1e2a33',
  'primaryTextColor':'#d6f5ff',
  'primaryBorderColor':'#5ccfe6',
  'lineColor':'#8a9199'
}}}%%
stateDiagram-v2
    [*] --> Typing : New() · active = screenTyping

    state Typing {
        [*] --> Loading
        Loading --> Active : quoteMsg
        Loading --> Failed : errMsg
        Active --> Active : KeyPressMsg / tickMsg
    }

    Typing --> Results : sessionFinishedMsg
    Results --> Typing : restartGameMsg (enter)
    Typing --> Typing : restartGameMsg (enter)

    Typing --> [*] : esc / ctrl+c
    Results --> [*] : esc / ctrl+c
```

