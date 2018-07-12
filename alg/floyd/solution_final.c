#include "stdio.h"
#include "stdlib.h"
#include "string.h"

#define VERTEX_NUM 10
#define EDGE_LIMIT 55
#define STR_LIMIT 32
#define INF 65535
#define MENU_SIZE 48

typedef unsigned short uint16;

typedef struct {
    char name[STR_LIMIT];
    char intro[STR_LIMIT];
} vertex;

typedef struct {
    uint16 source;
    uint16 target;
    uint16 weight;
} edge;

typedef struct {
    uint16 v_size;
    uint16 e_size;
    vertex v[VERTEX_NUM];
    uint16 adj[VERTEX_NUM][VERTEX_NUM];
} graph;

char *str_repeat(const char *string, int count)
{
    int i = 0;
    int j = 0;
    int size = 0;
    while (string[i] != '\0')
    {
        size++;
        i++;
    }
    char *ret = malloc((count*size+1)*sizeof(char));
    for (i=0; i<count; i++)
    {
        for (j=0; j<size; j++)
        {
            ret[i*size+j] = string[j];
        }
    }
    ret[size*count] = '\0';

    return ret;
}

graph* create_graph(const vertex *v, const uint16 v_size, const edge *e, const uint16 e_size)
{
    graph *g = (graph *)malloc(sizeof(graph));
    uint16 i, j;
    for (i=0; i<VERTEX_NUM; i++)
    {
        for (j=0; j<VERTEX_NUM; j++)
        {
            g->adj[i][j] = g->adj[j][i] = INF;
        }
    }
    g->v_size = v_size;
    g->e_size = e_size;
    for (i=0; i<v_size; i++)
    {
        g->v[i] = v[i];
    }
    for (i=0; i<e_size; i++)
    {
        uint16 source = e[i].source;
        uint16 target = e[i].target;
        uint16 weight = e[i].weight;
        g->adj[source][target] = g->adj[target][source] = weight;
    }

    return g;
}

void show_graph(graph *g)
{
    int i = 0;
    int j = 0;
    int v_size = (int)(g->v_size);
    char *line = str_repeat("-", MENU_SIZE);
    // char *tab = str_repeat(" ", 4);
    printf("%s\n", line);
    printf("Graph detail:\n");
    printf("%s\n", line);
    for (i=0; i<v_size; i++)
    {
        // printf("%s", tab);
        printf("Vertex Name: %s, Intro: %s\n", g->v[i].name, g->v[i].intro);
    }
    for (i=0; i<VERTEX_NUM; i++)
    {
        for (j=0; j<VERTEX_NUM; j++)
        {
            if (g->adj[i][j] != INF)
            {
                // printf("%s", tab);
                printf("Edge %d -> %d distance: %d\n", i, j, g->adj[i][j]);
            }
        }
    }
    printf("%s\n", line);
}

int** floyd(graph *g)
{
    int i, j, k;
    int **dp = (int **)malloc(g->v_size*sizeof(int*));
    for (i=0; i<g->v_size; i++)
    {
        dp[i] = (int *)malloc(g->v_size*sizeof(int));
        for (j=0; j<g->v_size; j++)
        {
            dp[i][j] = j;
            if (g->adj[i][j] == INF) {
                dp[i][j] = INF;
            }
        }
    }
    for (k=0; k<g->v_size; k++) 
    {
        for (i=0; i<g->v_size; i++) 
        {
            for (j=0; j<g->v_size; j++)
            {
                if(g->adj[i][k] + g->adj[k][j] < g->adj[i][j]) 
                {
                    g->adj[i][j] = g->adj[i][k] + g->adj[k][j];
                    dp[i][j] = dp[i][k];
                }
            }
        }
    }

    return dp;
}

void show_path(graph *g, int **path)
{
    char *line = str_repeat("-", MENU_SIZE);
    printf("%s\n", line);
    printf("Floyd result:\n");
    printf("%s\n", line);
    int i, j;
    for (i=0; i<g->v_size; i++)
    {
        for (j=0; j<g->v_size; j++)
        {
            if (path[i][j] != INF)
            {
                printf("Source %d to target %d next move is: %d\n", i, j, path[i][j]);
            }
        }
    }
    printf("%s\n", line);
}

void query_nd_show_shortest_path(int **path, uint16 source, uint16 target)
{
    printf("The shortest path of vertex %d to vextex %d is: ", source, target);
    uint16 limit = 0;
    while (source != target) {
        if (limit++ >= EDGE_LIMIT-1)
        {
            printf("detected infinite loop of routine");
            return;
        }
        if (path[source][target] == INF) {
            printf("there no way to ");
            break;
        } else {
            printf("%d->", source);
        }
        source = path[source][target];
    }
    printf("%d\n",target);
}

vertex *query_vertex(const graph *g, uint16 number)
{
    vertex *ret = (vertex *)(&g->v[number]);

    return ret;
}

char *get_vertex_name(const vertex *v)
{
    return (char *)(v->name);
}

char *get_vertex_intro(const vertex *v)
{
    return (char *)(v->intro);
}

void query_vertex_routine(graph *g)
{
    uint16 vertex_number;
    uint16 flag = 1;
    while(flag) 
    {
        printf("Give me the number of vertex:");
        scanf("%d", &vertex_number);
        if (vertex_number >= VERTEX_NUM)
        {
            printf("The number %d is too large(range [0, %d])!\n",vertex_number, VERTEX_NUM-1);
            continue;
        }
        printf("Vertex %d's name: %s\n", vertex_number, get_vertex_name(query_vertex(g, vertex_number)));
        printf("Vertex %d's intro: %s\n", vertex_number, get_vertex_intro(query_vertex(g, vertex_number)));
        flag = 0;
    }
}

void query_ssp_routine(int **path)
{
    uint16 src, tgt;
    uint16 flag = 1;
    while (flag)
    {
        printf("Give the source,target vertices number pair:");
        scanf("%d", &src);
        scanf("%d", &tgt);
        if (src >= VERTEX_NUM)
        {
            printf("The source number %d is too large(range [0, %d])!\n", src, VERTEX_NUM-1);
            continue;
        } else if (tgt >= VERTEX_NUM) {
            printf("The target number %d is too large(range [0, %d])!\n", tgt, VERTEX_NUM-1);
            continue;
        }
        printf("Query s(%d, %d)...\n", src, tgt);
        query_nd_show_shortest_path(path, src, tgt);
        flag = 0;
    }
}

int main()
{
    char name_repo[VERTEX_NUM][STR_LIMIT] = {
        "0","1","2","3","4","5","6","7","8","9"
    };
    char intro_repo[VERTEX_NUM][STR_LIMIT] = {
        "Apple",
        "Bnana",
        "Cow",
        "Diabolo",
        "Eggy",
        "Fuck",
        "Godness",
        "Holywood",
        "Independence",
        "Knowledge"
    };
    int i, j;
    vertex vertices[VERTEX_NUM];
    for (i=0; i<VERTEX_NUM; i++)
    {
        vertex v_tmp = *(vertex *)malloc(sizeof(vertex));
        strcpy(v_tmp.name, name_repo[i]);
        strcpy(v_tmp.intro, intro_repo[i]);
        vertices[i] = v_tmp;
    }
    edge edges[EDGE_LIMIT];
    int e_index = 0;
    for (i=0; i<VERTEX_NUM; i++)
    {
        for (j=0; j<VERTEX_NUM; j++)
        {
            if (i+j==VERTEX_NUM && e_index < EDGE_LIMIT-1)
            {
                edge e_tmp = *(edge *)malloc(sizeof(edge));
                e_tmp.source = i;
                e_tmp.target = j;
                e_tmp.weight = ((i*j)%VERTEX_NUM)+1;
                edges[e_index++] = e_tmp;
            }
        }
    }
    graph *g = create_graph(vertices, (uint16)VERTEX_NUM, edges, (uint16)e_index+1);
    show_graph(g);
    uint16 test_query_number = 0;
    vertex *vertex_0 = query_vertex(g, test_query_number);
    printf("Testing query Vertex %d's Name: %s\n", test_query_number, get_vertex_name(vertex_0));
    printf("Testing query Vertex %d's Intro: %s\n", test_query_number, get_vertex_intro(vertex_0));
    int **path = floyd(g);
    show_path(g, path);
    uint16 test_source = 2;
    uint16 test_target = 8;
    printf("Testing SSP problem s(%d, %d) now......\n", test_source, test_target);
    query_nd_show_shortest_path(path, 2, 8);
    uint16 choose;
    uint16 flag = 1;
    char *line = str_repeat("*", MENU_SIZE);
    while(flag) {
        printf("%s\n", line);
        printf("So, what do you want then?\n");
        printf("%s\n", line);
        printf("0. exit\n");
        printf("1. query vertext info\n");
        printf("2. query the shortest path between vertices\n");
        printf("%s\n", line);
        printf("Enter the number to make choise:");
        scanf("%d", &choose);
        switch (choose)
        {
            case 0:
                flag = 0;
                printf("Bye!\n");
                break;
            case 1:
                query_vertex_routine(g);
                break;
            case 2:
                query_ssp_routine(path);
                break;
        }
    }

    return 0;
}