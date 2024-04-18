import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebars: SidebarsConfig = {
  learnSidebar: [
    {
      type: "category",
      label: "Introduction",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "learn/introduction",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Background",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "learn/background",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Omni",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "learn/omni",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Testnet",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "learn/testnet",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "doc",
      id: "resources/glossary",
    },
  ],
  protocolSidebar: [
    {
      type: "category",
      label: "Introduction",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "protocol/introduction",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Security Model",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "protocol/security",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Cross-Rollup Messages",
      className: "sidebar-title",
      collapsible: false,
      items: [
        "protocol/xmessages/xmsg",
        {
          type: "category",
          label: "Components",
          className: "sidebar-title",
          collapsible: false,
          items: [
            "protocol/xmessages/components/portals",
            {
              type: "category",
              label: "Validators",
              collapsible: true,
              items: [
                {
                  type: "autogenerated",
                  dirName: "protocol/xmessages/components/validator",
                }
              ]
            },
            "protocol/xmessages/components/relayer",
          ]
        },
        "protocol/xmessages/fees",
        "protocol/xmessages/finality",
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "EVM Engine",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "protocol/evmengine",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    "protocol/audits",
    "protocol/future"
  ],
  developSidebar: [
    "develop/introduction",
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Testnet",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "develop/testnet",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "XApp",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "develop/xapp",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    "develop/contracts",
  ],
  operateSidebar: [
    {
      type: "category",
      label: "Introduction",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "operate/introduction",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Onboarding",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "operate/onboarding",
        }
      ]
    }
  ],
};

export default sidebars;
